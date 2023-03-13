package middleware

import (
	"net/http"
	"techku/Config"
)

type MiddlewareInterface interface {
	UserAuth(userPrivilege ...interface{}) func(http.Handler) http.Handler
}

type Middleware struct {
	MiddlewareInterface
	TokenHeader Config.TokenHeader
}
