package handlers

import (
	"encoding/json"
	"github.com/dayterr/gophkeeper_diploma/internal/authjwt"
	"github.com/dayterr/gophkeeper_diploma/internal/storage"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"strings"
)

func (ah *AsyncHandler) RegisterUser(w http.ResponseWriter, r *http.Request){
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("reading user data")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Info().Msg("error reading user data")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var u storage.User
	err = json.Unmarshal(body, &u)
	if err != nil {
		log.Info().Msg("error unmarshalling user data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if u.Login == "" || u.Password == "" {
		log.Info().Msg("empty fields recieved")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Info().Msg("encrypting user password")
	u.Password = authjwt.EncryptPassword(u.Password, ah.JWT_Key)

	log.Info().Msg("sending user data to database")
	err = ah.Storage.AddUser(r.Context(), u)
	if err != nil {
		log.Info().Msg("error saving user to database")
		if strings.Contains(err.Error(), "duplicate key") {
			w.WriteHeader(http.StatusConflict)
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token, err := authjwt.CreateToken(u.Login, ah.JWT_Key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "Bearer",
		Value: token,
	})
	w.WriteHeader(http.StatusOK)
}

func (ah *AsyncHandler) LogUser(w http.ResponseWriter, r *http.Request) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("reading user data")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Info().Msg("error reading user data")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	log.Info().Msg("unmarshalling user data")
	var u storage.User
	err = json.Unmarshal(body, &u)
	if err != nil {
		log.Info().Msg("error unmarshalling user data")
		w.WriteHeader(http.StatusInternalServerError)
	}

	if u.Login == "" || u.Password == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = ah.Storage.GetUser(r.Context(), u.Login)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Info().Msg("user successfully found")

	token, err := authjwt.CreateToken(u.Login, ah.JWT_Key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "Bearer",
		Value: token,
	})

	w.WriteHeader(http.StatusOK)
}
