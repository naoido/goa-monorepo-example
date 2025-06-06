package security

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	. "goa.design/goa/v3/dsl"
	"goa.design/goa/v3/security"
	"log"
	"os"
)

var secret []byte

func init() {
	secretStr := os.Getenv("JWT_SECRET")

	if secretStr == "" {
		log.Fatal("FATAL: JWT_SECRET environment variable is not set.")
	}

	secret = []byte(secretStr)
}

var JWTAuth = JWTSecurity("jwt", func() {
	Description("JWT")

	Scope("api:read", "APIリソースへの読み取りアクセス")
	Scope("api:write", "APIリソースへの書き込みアクセス")
	Scope("api:admin", "管理者アクセス")
})

func HasPermission(ctx context.Context, claims jwt.MapClaims, scheme *security.JWTScheme) (context.Context, error) {
	if claims == nil || claims["scopes"] == nil {
		return ctx, errors.New("missing scopes")
	}

	scopes, ok := claims["scopes"].([]interface{})
	if !ok {
		return ctx, errors.New("missing scopes")
	}

	scopesInToken := make([]string, len(scopes))
	for _, scope := range scopes {
		scopesInToken = append(scopesInToken, scope.(string))
	}

	if err := scheme.Validate(scopesInToken); err != nil {
		return ctx, err
	}

	ctx = contextWithAuthInfo(ctx, AuthInfo{Claims: claims})
	return ctx, nil
}

func ValidToken(tokenString string) (jwt.MapClaims, error) {
	claims := make(jwt.MapClaims)

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token has expired")
		}

		return nil, fmt.Errorf("invalid token: %w", err)
	}

	return claims, nil
}
