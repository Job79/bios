package auth

import (
	"bios/config"
	"bios/store"
	"encoding/base64"
	"net/http"
	"strings"
)

type Context struct {
	DB   store.Store
	Conf config.Config
}

// Authenticator is middleware that authenticates a request based on a token
// If the request is unauthenticated, it will abort the request and return a 401 status code
func (ctx Context) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get auth from header
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(auth) != 2 || auth[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Validate token
		tokenBytes, _ := base64.RawURLEncoding.DecodeString(auth[1])
		token, err := ctx.DB.FetchToken(tokenBytes, ctx.Conf.Security.Token)
		if err != nil || token.IsExpired() {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Continue request
		next.ServeHTTP(w, r)
	})
}
