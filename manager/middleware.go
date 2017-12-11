package manager

import (
	"context"
	"log"
	"net/http"
	"time"
)

type contextKey string

func (c contextKey) String() string {
	return "manager context key " + string(c)
}

// AddContext adds a golang context to the handler func which will be
// passed along the handlers
func AddContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AddUserTable adds the user table object to UserHandler contexts
func AddUserTable(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u := UserTable{
			db,
		}
		var userTable = contextKey("user-table")
		ctx := context.WithValue(r.Context(), userTable, &u)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Logger logs http(s) request to STDOUT
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
