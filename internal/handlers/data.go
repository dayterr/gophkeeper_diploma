package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/dayterr/gophkeeper_diploma/internal/authjwt"
	"github.com/dayterr/gophkeeper_diploma/internal/storage"
	"github.com/dayterr/gophkeeper_diploma/internal/validators"
)

func (ah *AsyncHandler) PostData(w http.ResponseWriter, r *http.Request) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("reading sent data")
	dataType := chi.URLParam(r, "dataType")
	if dataType == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Info().Msg("error reading data")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	switch dataType {
	case "cards":
		log.Info().Msg("card was sent")

		var c storage.Card

		err = json.Unmarshal(body, &c)
		if err != nil {
			log.Info().Msg("error unmarshalling data")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = validators.ValidateCard(c)
		if err != nil {
			log.Info().Msg("a field was empty")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		cyphered, err := authjwt.EncryptData(c.CardNumber, ah.CryptoKey)
		if err != nil {
			log.Info().Msg("error cyphering data")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		c.CardNumber = string(cyphered)

		cyphered, err = authjwt.EncryptData(c.ExpDate, ah.CryptoKey)
		if err != nil {
			log.Info().Msg("error cyphering data")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		c.ExpDate = string(cyphered)

		cyphered, err = authjwt.EncryptData(c.Cardholder, ah.CryptoKey)
		if err != nil {
			log.Info().Msg("error cyphering data")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		c.Cardholder = string(cyphered)

		cyphered, err = authjwt.EncryptData(c.CVV, ah.CryptoKey)
		if err != nil {
			log.Info().Msg("error cyphering data")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		c.CVV = string(cyphered)

		login := r.Context().Value("login").(string)

		err = ah.Storage.AddCard(r.Context(), login, c)
		if err != nil {
			log.Info().Msg("error saving card")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)

	case "passwords":
		log.Info().Msg("password was sent")

		var p storage.Password

		err = json.Unmarshal(body, &p)
		if err != nil {
			log.Info().Msg("error unmarshalling data")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = validators.ValidatePassword(p)
		if err != nil {
			log.Info().Msg("a field was empty")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		cyphered, err := authjwt.EncryptData(p.Login, ah.CryptoKey)
		if err != nil {
			log.Info().Msg("error cyphering data")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		p.Login = string(cyphered)

		cyphered, err = authjwt.EncryptData(p.Password, ah.CryptoKey)
		if err != nil {
			log.Info().Msg("error cyphering data")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		p.Password = string(cyphered)

		login := r.Context().Value("login").(string)

		err = ah.Storage.AddPassword(r.Context(), login, p)
		if err != nil {
			log.Info().Msg("error saving password")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)

	case "texts":
		log.Info().Msg("text was sent")

		var t storage.Text

		err = json.Unmarshal(body, &t)
		if err != nil {
			log.Info().Msg("error unmarshalling data")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = validators.ValidateText(t)
		if err != nil {
			log.Info().Msg("a field was empty")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		cyphered, err := authjwt.EncryptData(t.Data, ah.CryptoKey)
		if err != nil {
			log.Info().Msg("error cyphering data")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		t.Data = string(cyphered)

		login := r.Context().Value("login").(string)

		err = ah.Storage.AddText(r.Context(), login, t)
		if err != nil {
			log.Info().Msg("error saving text")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)

	case "files":
		log.Info().Msg("file was sent")

		var b storage.Binary

		err = json.Unmarshal(body, &b)
		if err != nil {
			log.Info().Msg("error unmarshalling data")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = validators.ValidateBinary(b)
		if err != nil {
			log.Info().Msg("a field was empty")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		cyphered, err := authjwt.EncryptData(b.Filename, ah.CryptoKey)
		if err != nil {
			log.Info().Msg("error cyphering data")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		b.Filename = string(cyphered)

		login := r.Context().Value("login").(string)

		err = ah.Storage.AddFile(r.Context(), login, b)
		if err != nil {
			log.Info().Msg("error saving file")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)

	default:
		w.WriteHeader(http.StatusBadRequest)
	}

}

func (ah *AsyncHandler) ListData(w http.ResponseWriter, r *http.Request) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("listing data r")
	dataType := chi.URLParam(r, "dataType")
	if dataType == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch dataType {
	case "cards":
		login := r.Context().Value("login").(string)

		cards, err := ah.Storage.ListCards(r.Context(), login)
		if err != nil {
			log.Info().Msg("error getting cards")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(cards) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		var cardsBody []storage.Card
		for _, card := range cards {
			decyphered, err := authjwt.DecryptData([]byte(card.CardNumber), ah.CryptoKey)
			if err != nil {
				log.Info().Msg("error decyphering data")
				w.WriteHeader(http.StatusInternalServerError)
			}
			card.CardNumber = string(decyphered)

			decyphered, err = authjwt.DecryptData([]byte(card.ExpDate), ah.CryptoKey)
			if err != nil {
				log.Info().Msg("error decyphering data")
				w.WriteHeader(http.StatusInternalServerError)
			}
			card.ExpDate = string(decyphered)

			decyphered, err = authjwt.DecryptData([]byte(card.Cardholder), ah.CryptoKey)
			if err != nil {
				log.Info().Msg("error decyphering data")
				w.WriteHeader(http.StatusInternalServerError)
			}
			card.Cardholder = string(decyphered)

			decyphered, err = authjwt.DecryptData([]byte(card.CVV), ah.CryptoKey)
			if err != nil {
				log.Info().Msg("error decyphering data")
				w.WriteHeader(http.StatusInternalServerError)
			}
			card.CVV = string(decyphered)
			cardsBody = append(cardsBody, card)
		}

		body, err := json.Marshal(&cardsBody)
		if err != err {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
		w.WriteHeader(http.StatusOK)

	case "passwords":
		login := r.Context().Value("login").(string)

		passwords, err := ah.Storage.ListPasswords(r.Context(), login)
		if err != nil {
			log.Info().Msg("error getting passwords")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(passwords) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		var passwordsBody []storage.Password
		for _, password := range passwords {
			decyphered, err := authjwt.DecryptData([]byte(password.Login), ah.CryptoKey)
			if err != nil {
				log.Info().Msg(err.Error())
				log.Info().Msg("error decyphering data")
				w.WriteHeader(http.StatusInternalServerError)
			}
			password.Login = string(decyphered)

			decyphered, err = authjwt.DecryptData([]byte(password.Password), ah.CryptoKey)
			if err != nil {
				log.Info().Msg(err.Error())
				log.Info().Msg("error decyphering data")
				w.WriteHeader(http.StatusInternalServerError)
			}
			password.Password = string(decyphered)

			passwordsBody = append(passwordsBody, password)
		}

		body, err := json.Marshal(&passwordsBody)
		if err != err {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
		w.WriteHeader(http.StatusOK)

	case "texts":
		login := r.Context().Value("login").(string)

		texts, err := ah.Storage.ListTexts(r.Context(), login)
		if err != nil {
			log.Info().Msg("error getting texts")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(texts) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		var textsBody []storage.Text
		for _, text := range texts {
			decyphered, err := authjwt.DecryptData([]byte(text.Data), ah.CryptoKey)
			if err != nil {
				log.Info().Msg("error decyphering data")
				w.WriteHeader(http.StatusInternalServerError)
			}
			text.Data = string(decyphered)

			textsBody = append(textsBody, text)
		}

		body, err := json.Marshal(&textsBody)
		if err != err {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
		w.WriteHeader(http.StatusOK)

	case "files":
		login := r.Context().Value("login").(string)

		files, err := ah.Storage.ListFiles(r.Context(), login)
		if err != nil {
			log.Info().Msg("error getting files")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(files) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		var filesBody []storage.Binary
		for _, binary := range files {
			decyphered, err := authjwt.DecryptData([]byte(binary.Filename), ah.CryptoKey)
			if err != nil {
				log.Info().Msg("error decyphering data")
				w.WriteHeader(http.StatusInternalServerError)
			}
			binary.Filename = string(decyphered)

			filesBody = append(filesBody, binary)
		}

		body, err := json.Marshal(&filesBody)
		if err != err {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
		w.WriteHeader(http.StatusOK)

	default:
		w.WriteHeader(http.StatusNotFound)
	}

}

func (ah *AsyncHandler) DeleteData(w http.ResponseWriter, r *http.Request) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("deleting a card")
	dataType := chi.URLParam(r, "dataType")
	if dataType == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	dataID := chi.URLParam(r, "dataID")
	if dataType == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch dataType {
	case "cards":
		cardID, err := strconv.Atoi(dataID)
		if err != nil {
			log.Info().Msg("error getting cardID")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		login := r.Context().Value("login").(string)

		err = ah.Storage.DeleteCard(r.Context(), int64(cardID), login)
		if err != nil {
			if strings.Contains(err.Error(), "this user can't") {
				log.Info().Msg("this card doesn't belong to this user")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err != nil {
			log.Info().Msg("error deleting the card from database")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	case "passwords":
		passwordID, err := strconv.Atoi(dataID)
		if err != nil {
			log.Info().Msg("error getting passwordID")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		login := r.Context().Value("login").(string)

		err = ah.Storage.DeletePassword(r.Context(), int64(passwordID), login)
		if err != nil {
			if strings.Contains(err.Error(), "this user can't") {
				log.Info().Msg("this password doesn't belong to this user")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err != nil {
			log.Info().Msg("error deleting the password from database")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	case "texts":
		textID, err := strconv.Atoi(dataID)
		if err != nil {
			log.Info().Msg("error getting textID")
			w.WriteHeader(http.StatusInternalServerError)
		}

		login := r.Context().Value("login").(string)

		err = ah.Storage.DeleteText(r.Context(), int64(textID), login)
		if err != nil {
			if strings.Contains(err.Error(), "this user can't") {
				log.Info().Msg("this text doesn't belong to this user")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err != nil {
			log.Info().Msg("error deleting the text from database")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	case "files":
		fileID, err := strconv.Atoi(dataID)
		if err != nil {
			log.Info().Msg("error getting fileID")
			w.WriteHeader(http.StatusInternalServerError)
		}

		login := r.Context().Value("login").(string)

		err = ah.Storage.DeleteFile(r.Context(), int64(fileID), login)
		if err != nil {
			if strings.Contains(err.Error(), "this user can't") {
				log.Info().Msg("this file doesn't belong to this user")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err != nil {
			log.Info().Msg("error deleting the file from database")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
