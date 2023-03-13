package middleware

import (
	"net/http"
)

func (m Middleware) UserAuth(userPrivilege ...interface{}) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			//Helper.HttpResponseErrorNew(w, r, nil, Constant.StatusUnauthorizedInvalidToken)
			return
		})
	}
}
