package tuiagent

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"

	"github.com/rivo/tview"

	"github.com/dayterr/gophkeeper_diploma/internal/storage"
	"github.com/dayterr/gophkeeper_diploma/internal/validators"
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
				Certificates:       []tls.Certificate{cert},
			},
		}}

	t.Address = addr
	t.TUIApp = tview.NewApplication()

	return t, nil
}

func (t TUIClient) SendCard(card storage.Card) error {
	url := fmt.Sprintf("https://%v/cards/", t.Address)

	cardJSON, err := json.Marshal(card)
	if err != nil {
		return err
	}

	resp, err := t.HTTPSender.Post(url, "application/json", bytes.NewBuffer(cardJSON))

	if resp.StatusCode == http.StatusBadRequest {
		return validators.ErrorCodeBadRequestRecieved
	}

	if resp.StatusCode == http.StatusInternalServerError {
		return validators.ErrorCodeInternalErrorRecieved
	}

	if err != nil {
		return err
	}

	return nil
}

func (t TUIClient) ListCards() ([]storage.Card, error) {
	url := fmt.Sprintf("https://%v/cards/", t.Address)

	resp, err := t.HTTPSender.Get(url)
	if err != nil {
		return []storage.Card{}, err
	}

	var cards []storage.Card
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &cards)

	return cards, nil
}

func (t TUIClient) DeleteCard(cardID string) error {
	url := fmt.Sprintf("https://%v/cards/%v", t.Address, cardID)
	req, err := http.NewRequest(
		http.MethodDelete, url, nil,
	)
	if err != nil {
		return err
	}

	resp, err := t.HTTPSender.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return validators.ErrorDeleteFailed
	}

	return nil
}

func (t TUIClient) ListPasswords() ([]storage.Password, error) {
	url := fmt.Sprintf("https://%v/passwords/", t.Address)

	resp, err := t.HTTPSender.Get(url)
	if err != nil {
		return []storage.Password{}, err
	}

	var passwords []storage.Password
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &passwords)

	return passwords, nil
}

func (t TUIClient) SendPassword(password storage.Password) error {
	url := fmt.Sprintf("https://%v/passwords/", t.Address)

	passwordJSON, err := json.Marshal(password)
	if err != nil {
		return err
	}

	resp, err := t.HTTPSender.Post(url, "application/json", bytes.NewBuffer(passwordJSON))

	if resp.StatusCode == http.StatusBadRequest {
		return validators.ErrorCodeBadRequestRecieved
	}

	if resp.StatusCode == http.StatusInternalServerError {
		return validators.ErrorCodeInternalErrorRecieved
	}

	if err != nil {
		return err
	}

	return nil
}

func (t TUIClient) DeletePassword(cardID string) error {
	url := fmt.Sprintf("https://%v/passwords/%v", t.Address, cardID)
	req, err := http.NewRequest(
		http.MethodDelete, url, nil,
	)
	if err != nil {
		return err
	}

	resp, err := t.HTTPSender.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return validators.ErrorDeleteFailed
	}

	return nil
}

func (t TUIClient) SendText(text storage.Text) error {
	url := fmt.Sprintf("https://%v/texts/", t.Address)

	textJSON, err := json.Marshal(text)
	if err != nil {
		return err
	}

	resp, err := t.HTTPSender.Post(url, "application/json", bytes.NewBuffer(textJSON))

	if resp.StatusCode == http.StatusBadRequest {
		return validators.ErrorCodeBadRequestRecieved
	}

	if resp.StatusCode == http.StatusInternalServerError {
		return validators.ErrorCodeInternalErrorRecieved
	}

	if err != nil {
		return err
	}

	return nil
}

func (t TUIClient) ListTexts() ([]storage.Text, error) {
	url := fmt.Sprintf("https://%v/texts/", t.Address)

	resp, err := t.HTTPSender.Get(url)
	if err != nil {
		return []storage.Text{}, err
	}

	var texts []storage.Text
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &texts)

	return texts, nil
}

func (t TUIClient) DeleteText(textID string) error {
	url := fmt.Sprintf("https://%v/texts/%v", t.Address, textID)
	req, err := http.NewRequest(
		http.MethodDelete, url, nil,
	)
	if err != nil {
		return err
	}

	resp, err := t.HTTPSender.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return validators.ErrorDeleteFailed
	}

	return nil
}

func (t TUIClient) SendFile(binary storage.Binary) error {
	url := fmt.Sprintf("https://%v/files/", t.Address)

	fileJSON, err := json.Marshal(binary)
	if err != nil {
		return err
	}

	resp, err := t.HTTPSender.Post(url, "application/json", bytes.NewBuffer(fileJSON))

	if resp.StatusCode == http.StatusBadRequest {
		return validators.ErrorCodeBadRequestRecieved
	}

	if resp.StatusCode == http.StatusInternalServerError {
		return validators.ErrorCodeInternalErrorRecieved
	}

	if err != nil {
		return err
	}

	return nil
}

func (t TUIClient) ListFiles() ([]storage.Binary, error) {
	url := fmt.Sprintf("https://%v/files/", t.Address)

	resp, err := t.HTTPSender.Get(url)
	if err != nil {
		return []storage.Binary{}, err
	}

	var files []storage.Binary
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &files)

	return files, nil
}

func (t TUIClient) DeleteFile(fileID string) error {
	url := fmt.Sprintf("https://%v/files/%v", t.Address, fileID)
	req, err := http.NewRequest(
		http.MethodDelete, url, nil,
	)
	if err != nil {
		return err
	}

	resp, err := t.HTTPSender.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return validators.ErrorDeleteFailed
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
		return validators.ErrorAlreadyRegistered
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
		return validators.ErrorLoginNotFound
	}

	if err != nil {
		return err
	}

	return nil
}
