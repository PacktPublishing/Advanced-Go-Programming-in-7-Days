package handler

import (
	"net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-6/5-Authentication/easy-issues/web"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-6/5-Authentication/easy-issues/application"
	"errors"
)

func JWTAuthHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")
		authArr := strings.Split(authToken, " ")

		if len(authArr) != 2 {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
		}

		jwtToken := authArr[1]

		token, err := jwt.ParseWithClaims(jwtToken, &web.JWTData{}, func(token *jwt.Token) (interface{}, error) {
			if jwt.SigningMethodHS256 != token.Method {
				return nil, errors.New("invalid signing algorithm")
			}

			return []byte(application.Secret), nil
		})

		if err != nil {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		if token.Valid {
			h(w, r)
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
		}
	}
}


