package greet

import (
	"context"
	"fmt"
	greet "goa-example/microservices/greet/gen/greet"
	security2 "goa-example/pkg/security"

	"goa.design/clue/log"
	"goa.design/goa/v3/security"
)

// greet service example implementation.
// The example methods log the requests and return zero values.
type greetsrvc struct{}

// NewGreet returns the greet service implementation.
func NewGreet() greet.Service {
	return &greetsrvc{}
}

// JWTAuth implements the authorization logic for service "greet" for the "jwt"
// security scheme.
func (s *greetsrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	claims, err := security2.ValidToken(token)
	if err != nil {
		return ctx, err
	}

	return security2.HasPermission(ctx, claims, scheme)
}

// Greet method
func (s *greetsrvc) Greet(ctx context.Context) (res string, err error) {
	return "Hello World!", nil
}

// Hello method
func (s *greetsrvc) Hello(ctx context.Context, p *greet.HelloPayload) (res string, err error) {
	username := security2.ContextAuthInfo(ctx).Claims["sub"].(string)
	log.Print(ctx, log.KV{K: "username", V: username})

	return fmt.Sprintf("Hello %s!", username), nil
}
