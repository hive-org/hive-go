package hive_go

import (
	"net/http"
)

func (client *Client) middleware(next http.Handler, required bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		status, user := client.Login(ctx, r)
		if status == Ok && user != nil || !required {
			ctx = client.SetUserToContext(ctx, user)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}

func (client *Client) AuthRequired(handler http.Handler) http.Handler {
	return client.middleware(handler, true)
}

func (client *Client) Auth(handler http.Handler) http.Handler {
	return client.middleware(handler, false)
}
