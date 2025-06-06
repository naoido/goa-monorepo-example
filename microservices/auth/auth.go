package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"goa-example/microservices/auth/gen/auth"
	"goa-example/pkg/security"
	security2 "goa.design/goa/v3/security"
	"golang.org/x/net/context"
	"log"
	"os"
	"time"
)

var (
	accessSecret      []byte
	inMemoryBlackList = make([]string, 0)
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func init() {
	accessSecretStr := os.Getenv("JWT_SECRET")

	if accessSecretStr == "" {
		log.Fatal("FATAL: JWT_SECRET environment variable is not set.")
	}

	accessSecret = []byte(accessSecretStr)
}

func (*AuthService) JWTAuth(ctx context.Context, token string, scheme *security2.JWTScheme) (context.Context, error) {
	claims, err := security.ValidToken(token)
	if err != nil {
		return ctx, err
	}

	return security.HasPermission(ctx, claims, scheme)
}

func (a *AuthService) Login(ctx context.Context, p *auth.LoginPayload) (*auth.LoginResult, error) {
	if p.Username == "" || p.Password == "" {
		return nil, auth.InvalidArgument("invalid argument")
	}

	if p.Username != "admin" || p.Password != "password" {
		return nil, auth.InvalidArgument("invalid username or password")
	}

	singedAccessToken, err := signToken(getAccessTokenClaims(p.Username))
	singedRefreshToken, err := signToken(getRefreshTokenClaims(p.Username))

	if err != nil {
		return nil, auth.Internal("failed to sign token")
	}

	res := &auth.LoginResult{
		AccessToken:  singedAccessToken,
		RefreshToken: singedRefreshToken,
	}

	return res, nil
}

func (a *AuthService) Refresh(ctx context.Context, p *auth.RefreshPayload) (*auth.RefreshResult, error) {
	if p.RefreshToken == "" {
		return nil, auth.InvalidArgument("invalid argument")
	}

	log.Println(p.RefreshToken)
	token, err := security.ValidToken(p.RefreshToken)
	log.Println(err)
	if err != nil {
		return nil, auth.InvalidArgument("invalid refresh token")
	}

	username, err := token.GetSubject()
	if err != nil {
		return nil, auth.InvalidArgument("invalid refresh token")
	}

	singedAccessToken, err := signToken(getAccessTokenClaims(username))
	singedRefreshToken, err := signToken(getRefreshTokenClaims(username))

	if err != nil {
		return nil, auth.Internal("failed to sign token")
	}

	res := &auth.RefreshResult{
		AccessToken:  singedAccessToken,
		RefreshToken: singedRefreshToken,
	}

	return res, nil
}

func (a *AuthService) Logout(ctx context.Context, p *auth.LogoutPayload) error {
	inMemoryBlackList = append(inMemoryBlackList, security.ContextAuthInfo(ctx).Claims["id"].(string))
	return nil
}

func getAccessTokenClaims(username string) jwt.Claims {
	return jwt.MapClaims{
		"id":     uuid.NewString(),
		"sub":    username,
		"iat":    jwt.NewNumericDate(time.Now()),
		"exp":    jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		"iss":    "GoaAuthService",
		"scopes": []string{"api:admin", "api:read", "api:write"},
	}
}

func getRefreshTokenClaims(username string) jwt.Claims {
	return jwt.MapClaims{
		"id":  uuid.NewString(),
		"sub": username,
		"iat": jwt.NewNumericDate(time.Now()),
		"exp": jwt.NewNumericDate(time.Now().AddDate(0, 0, 90)),
		"iss": "GoaAuthService",
	}
}

func signToken(claims jwt.Claims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(accessSecret)
}
