package storage

import (
	"context"
	"database/sql"
)

type Storager interface {
	AddUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, login string) (int64, error)
	AddCard(ctx context.Context, userID int64, card Card) error
	ListCards(ctx context.Context, userID int64) ([]Card, error)
	DeleteCard(ctx context.Context, userID, cardID int64) error
	AddPassword(ctx context.Context, userID int64, password Password) error
	ListPasswords(ctx context.Context, userID int64) ([]Password, error)
	DeletePassword(ctx context.Context, userID, passwordID int64) error
	AddText(ctx context.Context, userID int64, text Text) error
	ListTexts(ctx context.Context, userID int64) ([]Text, error)
	DeleteText(ctx context.Context, userID, textID int64) error
	AddFile(ctx context.Context, userID int64, binary Binary) error
	ListFiles(ctx context.Context, userID int64) ([]Binary, error)
	DeleteFile(ctx context.Context, userID, fileID int64) error
}

type DBStorage struct {
	DB *sql.DB
	DSN string
}

type User struct {
	ID int
	Login string `json:"login"`
	Password string `json:"password"`
}

type Card struct {
	ID int
	CardNumber string `json:"card_number"`
	ExpDate string `json:"exp_date"`
	Cardholder string `json:"cardholder"`
	CVV string `json:"cvv"`
	Metadata string `json:"metadata"`
}

type Password struct {
	ID int
	Login string `json:"login"`
	Password string `json:"password"`
	Metadata string `json:"metadata"`
}

type Text struct {
	ID int
	Data string `json:"data"`
	Metadata string `json:"metadata"`
}

type Binary struct {
	ID int
	Filename string `json:"filename"`
	Data []byte `json:"data"`
	Metadata string `json:"metadata"`
}