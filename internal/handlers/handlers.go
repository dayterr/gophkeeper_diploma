package handlers

import (
	"context"
	"github.com/dayterr/gophkeeper_diploma/internal/authjwt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"

	"github.com/dayterr/gophkeeper_diploma/internal/storage"
)

func NewAsyncHandler(dsn, jwtKey, cryptoKey string) (*AsyncHandler, error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("creating new async handler")

	dbStorage, err := storage.NewDB(dsn)
	if err != nil {
		return &AsyncHandler{}, err
	}

	ah := AsyncHandler{Storage: dbStorage, JWT_Key: jwtKey, CryptoKey: cryptoKey}
	log.Info().Msg("handler created successfully")

	return &ah, nil
}

func (ah *AsyncHandler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urls := []string{"/users/register", "/users/login"}
		path := r.URL.Path
		for _, v := range urls {
			if v == path {
				next.ServeHTTP(w, r)
				return
			}
		}

		cookieToken, err := r.Cookie("Bearer")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims := &authjwt.CustomClaims{}
		token, err := jwt.ParseWithClaims(cookieToken.Value, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(ah.JWT_Key), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "login", claims.Login)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

	})
}