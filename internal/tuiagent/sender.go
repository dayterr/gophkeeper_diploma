package tuiagent

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/dayterr/gophkeeper_diploma/internal/storage"
	"github.com/rivo/tview"
	"net/http"
	"net/http/cookiejar"
)

func NewTUICLient(certPath, certKeyPath, addr string) (TUIClient, error) {
	var t TUIClient

	cert, err := tls.LoadX509KeyPair(certPath, certKeyPath)
	if err != nil {
		return TUIClient{}, err
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return TUIClient{}, err
	}

	t.HTTPSender = &http.Client{Jar: jar,
		Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			Certificates: []tls.Certificate{cert},
		},
		}}

	t.Address = addr
	t.TUIApp = tview.NewApplication()

	return t, nil
}

func (t TUIClient) SendCard(card storage.Card) error {
	url := fmt.Sprintf("http://%v/cards/", t.Address)

	cardJSON, err := json.Marshal(card)
	if err != nil {
		return err
	}

	_, err = t.HTTPSender.Post(url, "application/json", bytes.NewBuffer(cardJSON))
	if err != nil {
		return err
	}

	return nil
}

func (t TUIClient) RegisterUser(user storage.User) error {
	url := fmt.Sprintf("https://%v/users/register", t.Address)

	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	r, err := t.HTTPSender.Post(url, "application/json", bytes.NewBuffer(userJSON))

	if r.StatusCode == http.StatusConflict {
		return ErrorAlreadyRegistered
	}

	if err != nil {
		return err
	}

	return nil
}

func (t TUIClient) LogUser(user storage.User) error {
	url := fmt.Sprintf("https://%v/users/login", t.Address)

	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	r, err := t.HTTPSender.Post(url, "application/json", bytes.NewBuffer(userJSON))

	if r.StatusCode == http.StatusNotFound {
		return ErrorLoginNotFound
	}

	if err != nil {
		return err
	}

	return nil
}