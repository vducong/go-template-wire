package middleware

import (
	"github.com/google/wire"
)

type AuthMiddlewares struct {
	Internal *InternalAuthMiddleware
	JWT      *JWTAuthMiddleware
}

var AuthMiddlewareSet = wire.NewSet(
	NewInternalAuthMiddleware,
	NewJWTAuthMiddleware,
	wire.Struct(new(AuthMiddlewares), "Internal", "JWT"),
)
