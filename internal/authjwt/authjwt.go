package authjwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func CreateToken(login, key string) (string, error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("creating new jwt token")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomClaims{
		Login: login,
		StandardClaims: jwt.StandardClaims {
			ExpiresAt: jwt.At(time.Now().Add(2 * time.Minute)),
			IssuedAt:  jwt.At(time.Now()),
		},
	})
	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	log.Info().Msg("token created successfully")

	return signedToken, nil
}

func EncryptPassword(password string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}