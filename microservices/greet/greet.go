package greet

import (
	"fmt"
	"goa-example/microservices/greet/gen/greet"
	"goa-example/pkg/security"
	security2 "goa.design/goa/v3/security"
	"golang.org/x/net/context"
)

type GreetService struct{}

func NewGreetService() *GreetService {
	return &GreetService{}
}

func (*GreetService) JWTAuth(ctx context.Context, token string, scheme *security2.JWTScheme) (context.Context, error) {
	claims, err := security.ValidToken(token)
	if err != nil {
		return ctx, err
	}

	return security.HasPermission(ctx, claims, scheme)
}

func (*GreetService) Greet(ctx context.Context) (string, error) {
	return "Hello World!", nil
}

func (*GreetService) Hello(ctx context.Context, p *greet.HelloPayload) (string, error) {
	username := security.ContextAuthInfo(ctx).Claims["sub"].(string)

	return fmt.Sprintf("Hello %s!", username), nil
}
