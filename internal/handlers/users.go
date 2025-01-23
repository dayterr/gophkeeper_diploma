package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/dayterr/gophkeeper_diploma/internal/authjwt"
	"github.com/dayterr/gophkeeper_diploma/internal/storage"
	"github.com/dayterr/gophkeeper_diploma/internal/validators"
)

func (ah *AsyncHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
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

	err = validators.ValidateUser(u)
	if err != nil {
		log.Info().Msg("a field was invalid")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
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

		var pqErr *pq.Error
		ok := errors.As(err, &pqErr)
		if ok && pqErr.Code == storage.DupErr {
			w.WriteHeader(http.StatusConflict)
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token, err := authjwt.CreateToken(u.Login, ah.JWT_Key)
	if err != nil {
		log.Info().Msg("error creating token: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "Bearer",
		Value: token,
		Path:  "/",
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
		log.Info().Msg("login and password fields can't be empty")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = ah.Storage.GetUser(r.Context(), u.Login)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			log.Info().Msg("user wasn't found in the database" + err.Error())
			w.WriteHeader(http.StatusNotFound)
			return
		}

		log.Info().Msg("error getting user from database occurred: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Info().Msg("user successfully found")

	token, err := authjwt.CreateToken(u.Login, ah.JWT_Key)
	if err != nil {
		log.Info().Msg("error creating token: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "Bearer",
		Value: token,
		Path:  "/",
	})

	w.WriteHeader(http.StatusOK)
}
