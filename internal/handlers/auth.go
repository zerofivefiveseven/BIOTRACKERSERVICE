package handlers

import (
	"BIOTRACKERSERVICE/internal/auth"
	"errors"
	"net/http"
	"strings"
)

func (c *Controller) UserAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		tokenSplit := strings.Fields(token)

		if len(tokenSplit) < 2 || tokenSplit[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token = tokenSplit[1]
		//TODO: errors from different package
		if token == "" {
			ProcessError(w, ErrNoToken, http.StatusBadRequest)
			return
		}
		err := c.Usecases.UserAuth(r.Context(), token)
		if errors.Is(err, auth.ErrInvalidToken) || errors.Is(err, auth.ErrForbidden) {
			ProcessError(w, err, http.StatusForbidden)
			return
		}
		if errors.Is(err, auth.ErrUnauthorized) {
			ProcessError(w, err, http.StatusUnauthorized)
			return
		}
		if err != nil {
			ProcessError(w, err, http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	})
}
