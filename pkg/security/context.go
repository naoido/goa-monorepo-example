package security

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
)

type AuthInfo struct {
	Claims jwt.MapClaims
}

func (auth AuthInfo) String() string {
	if auth.Claims == nil {
		return "AuthInfo: UnAuthenticated"
	}

	return "AuthInfo: Authenticated"
}

type ctxValue int

const (
	ctxValueClaims ctxValue = iota
)

func contextWithAuthInfo(ctx context.Context, auth AuthInfo) context.Context {
	return context.WithValue(ctx, ctxValueClaims, auth)
}

func ContextAuthInfo(ctx context.Context) (auth AuthInfo) {
	auth, _ = ctx.Value(ctxValueClaims).(AuthInfo)
	return
}
