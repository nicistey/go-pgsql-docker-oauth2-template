//middleware.go
package api

import (
 "context"
 "log"
 "net/http"
 "strings"

 "github.com/golang-jwt/jwt/v5"
)

func (api *api) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, r.Method)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	
		if r.Method == "OPTIONS" {
			return
		}
	
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			if (strings.HasPrefix(r.URL.Path,"/api/events"))||(strings.HasPrefix(r.URL.Path,"/auth")) || (strings.HasPrefix(r.URL.Path,"/health")){ //Проверка только для /api/events
				next.ServeHTTP(w, r)
				return
			}
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(api.cfg.SecretJWTKey), nil
		})
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
	
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "userID", claims.ID) //Добавляем ID пользователя в контекст
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
		}
	})
}
   