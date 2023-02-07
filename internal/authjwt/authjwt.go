package authjwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"time"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"errors"
	"fmt"

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
			ExpiresAt: jwt.At(time.Now().Add(1440 * time.Minute)),
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

func EncryptData(data string, key string) ([]byte, error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("starting cyphering data")

	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Info().Msg(key)
		log.Info().Msg(err.Error())
		log.Info().Msg("error creating a new cypher")
		return []byte{}, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Info().Msg("error creating a new block")
		return []byte{}, err
	}

	nonce := make([]byte, gcm.NonceSize())
	fmt.Println([]byte(data))
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		log.Info().Msg("error reading nonce")
		return []byte{}, err
	}
	
	return gcm.Seal(nonce, nonce, []byte(data), nil), nil
}

func DecryptData(data []byte, key string) ([]byte, error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("starting decyphering data")

	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Info().Msg("error creating a new decypher")
		return []byte{}, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Info().Msg("error creating a new block")
		return []byte{}, err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return []byte{}, errors.New("ciphertext too short")
	}
	
	nonce, data := data[:nonceSize], data[nonceSize:]

	log.Info().Msg("data decyphered successfuly")

	return gcm.Open(nil, nonce, data, nil)
}