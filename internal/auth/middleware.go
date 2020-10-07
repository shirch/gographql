package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/shirch/graphql/graph/model"
	"github.com/shirch/graphql/internal/pkg/jwt"
	"github.com/shirch/graphql/internal/users"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			query := r.URL.Query()
			fmt.Printf("query: %v", query)
			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// create user and check if user exists in db
			user := model.User{Name: username}
			id, err := users.GetUserIdByUsername(username, db)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			user.ID = id
			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}
