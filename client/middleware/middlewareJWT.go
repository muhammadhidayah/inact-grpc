package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var APPLICATION_NAME = "INACT Cloud"
var LOGIN_EXPIRATION_DURATION = time.Duration(3) * time.Hour
var COOKIE_EXPIRATION_DURATION = time.Duration(2) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("s3cr337 0f 1n4c7")

type MyClaims struct {
	jwt.StandardClaims
	PersonNo string `json:"id"`
	PersonID string `json:"username"`
}

func MiddlewareJWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/auth/login" {
			next.ServeHTTP(w, r)
			return
		}

		var tokenString string
		storedCookie, _ := r.Cookie("RefreshTokenJWT")
		if storedCookie != nil {
			tokenString = storedCookie.Value
		} else {
			authHeader := r.Header.Get("Authorization")
			if !strings.Contains(authHeader, "Bearer") {
				http.Error(w, "Invalid token", http.StatusBadRequest)
				return
			}

			tokenString = strings.Replace(authHeader, "Bearer ", "", -1)
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signing method invalid")
			} else if method != JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("Signing method invalid")
			}

			return JWT_SIGNATURE_KEY, nil
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(context.Background(), "userInfo", claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
