package handlers

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"encoding/json"
	"strconv"
	"github.com/dayterr/gophkeeper_diploma/internal/storage"
	"github.com/dayterr/gophkeeper_diploma/internal/authjwt"
	"github.com/dayterr/gophkeeper_diploma/internal/validators"
	"strings"
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

		userLogin := r.Context().Value("login").(string)
		userID, err := ah.Storage.GetUser(r.Context(), userLogin)
		if err != nil {
			log.Info().Msg("couldn't find the user")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		err = ah.Storage.AddCard(r.Context(), userID, c)
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

		userLogin := r.Context().Value("login").(string)
		userID, err := ah.Storage.GetUser(r.Context(), userLogin)
		if err != nil {
			log.Info().Msg("couldn't find the user")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		err = ah.Storage.AddPassword(r.Context(), userID, p)
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

		userLogin := r.Context().Value("login").(string)
		userID, err := ah.Storage.GetUser(r.Context(), userLogin)
		if err != nil {
			log.Info().Msg("couldn't find the user")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		err = ah.Storage.AddText(r.Context(), userID, t)
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

		userLogin := r.Context().Value("login").(string)
		userID, err := ah.Storage.GetUser(r.Context(), userLogin)
		if err != nil {
			log.Info().Msg("couldn't find the user")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		err = ah.Storage.AddFile(r.Context(), userID, b)
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
		userLogin := r.Context().Value("login").(string)
		userID, err := ah.Storage.GetUser(r.Context(), userLogin)
		if err != nil {
			log.Info().Msg("couldn't find the user")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		cards, err := ah.Storage.ListCards(r.Context(), userID)
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
		userLogin := r.Context().Value("login").(string)
		userID, err := ah.Storage.GetUser(r.Context(), userLogin)
		if err != nil {
			log.Info().Msg("couldn't find the user")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		passwords, err := ah.Storage.ListPasswords(r.Context(), userID)
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
		userLogin := r.Context().Value("login").(string)
		userID, err := ah.Storage.GetUser(r.Context(), userLogin)
		if err != nil {
			log.Info().Msg("couldn't find the user")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		texts, err := ah.Storage.ListTexts(r.Context(), userID)
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
		userLogin := r.Context().Value("login").(string)
		userID, err := ah.Storage.GetUser(r.Context(), userLogin)
		if err != nil {
			log.Info().Msg("couldn't find the user")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		files, err := ah.Storage.ListFiles(r.Context(), userID)
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

		userLogin := r.Context().Value("login").(string)
		userID, err := ah.Storage.GetUser(r.Context(), userLogin)
		if err != nil {
			log.Info().Msg("couldn't find the user")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		
		err = ah.Storage.DeleteCard(r.Context(), userID, int64(cardID))
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

		userLogin := r.Context().Value("login").(string)
		userID, err := ah.Storage.GetUser(r.Context(), userLogin)
		if err != nil {
			log.Info().Msg("couldn't find the user")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		err = ah.Storage.DeletePassword(r.Context(), userID, int64(passwordID))
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

		userLogin := r.Context().Value("login").(string)
		userID, err := ah.Storage.GetUser(r.Context(), userLogin)
		if err != nil {
			log.Info().Msg("couldn't find the user")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		err = ah.Storage.DeleteText(r.Context(), userID, int64(textID))
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

		userLogin := r.Context().Value("login").(string)
		userID, err := ah.Storage.GetUser(r.Context(), userLogin)
		if err != nil {
			log.Info().Msg("couldn't find the user")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		err = ah.Storage.DeleteFile(r.Context(), userID, int64(fileID))
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