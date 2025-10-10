package middleware

import (
	user_memory_cache "formaura/pkg/cache/user_memory"
	"formaura/pkg/constants"
	user_repo "formaura/pkg/repositories/user"

	"context"
	"formaura/pkg/output"
	"formaura/pkg/services/jwt"
	"formaura/pkg/services/uuid"
	"net/http"
)

func AuthCachedMiddleware(repo user_repo.Repository, cache *user_memory_cache.Cache) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get(constants.AUTH_TOKEN_HEADER)
			if len(token) == 0 {
				output.WriteJson(w, r, http.StatusForbidden, output.MessageResponse{Message: "Auth token required"})
				return
			}

			parsed, err := jwt.Parse(token)
			if err != nil {
				output.WriteJson(w, r, http.StatusForbidden, output.MessageResponse{Message: err.Error()})
				return
			}

			id, ok := parsed["uuid"].(string)
			if !ok || !uuid.Validate(id) {
				output.WriteJson(w, r, http.StatusForbidden, output.MessageResponse{Message: "Auth token invalid"})
				return
			}

			if usr := cache.Get(id); usr != nil {
				ctx := context.WithValue(r.Context(), constants.USER_CTX, usr)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			usr, err := repo.GetByUUID(r.Context(), id)
			if err != nil || usr == nil {
				output.WriteJson(w, r, http.StatusForbidden, output.MessageResponse{Message: "Auth failed"})
				return
			}

			cache.Set(id, usr)

			ctx := context.WithValue(r.Context(), constants.USER_CTX, usr)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func AuthAlwaysFreshMiddleware(repo user_repo.Repository, cache *user_memory_cache.Cache) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get(constants.AUTH_TOKEN_HEADER)
			if len(token) == 0 {
				output.WriteJson(w, r, http.StatusForbidden, output.MessageResponse{Message: "Auth token required"})
				return
			}

			parsed, err := jwt.Parse(token)
			if err != nil {
				output.WriteJson(w, r, http.StatusForbidden, output.MessageResponse{Message: err.Error()})
				return
			}

			id, ok := parsed["uuid"].(string)
			if !ok || !uuid.Validate(id) {
				output.WriteJson(w, r, http.StatusForbidden, output.MessageResponse{Message: "Auth token invalid"})
				return
			}

			usr, err := repo.GetByUUID(r.Context(), id)
			if err != nil || usr == nil {
				output.WriteJson(w, r, http.StatusForbidden, output.MessageResponse{Message: "Auth failed"})
				return
			}

			cache.Set(id, usr)

			ctx := context.WithValue(r.Context(), constants.USER_CTX, usr)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
