package storage

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type Storager interface {
	AddUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, login string) (int64, error)
	AddCard(ctx context.Context, login string, card Card) error
	ListCards(ctx context.Context, login string) ([]Card, error)
	DeleteCard(ctx context.Context, cardID int64, login string) error
	AddPassword(ctx context.Context, login string, password Password) error
	ListPasswords(ctx context.Context, login string) ([]Password, error)
	DeletePassword(ctx context.Context, passwordID int64, login string) error
	AddText(ctx context.Context, login string, text Text) error
	ListTexts(ctx context.Context, login string) ([]Text, error)
	DeleteText(ctx context.Context, textID int64, login string) error
	AddFile(ctx context.Context, login string, binary Binary) error
	ListFiles(ctx context.Context, login string) ([]Binary, error)
	DeleteFile(ctx context.Context, fileID int64, login string) error
}

type DBStorage struct {
	DB  *sql.DB
	DSN string
}

type User struct {
	ID       int
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Card struct {
	ID         int
	CardNumber string `json:"card_number"`
	ExpDate    string `json:"exp_date"`
	Cardholder string `json:"cardholder"`
	CVV        string `json:"cvv"`
	Metadata   string `json:"metadata"`
}

type Password struct {
	ID       int
	Login    string `json:"login"`
	Password string `json:"password"`
	Metadata string `json:"metadata"`
}

type Text struct {
	ID       int
	Data     string `json:"data"`
	Metadata string `json:"metadata"`
}

type Binary struct {
	ID       int
	Filename string `json:"filename"`
	Data     []byte `json:"data"`
	Metadata string `json:"metadata"`
}

const (
	DupErr = pq.ErrorCode("23505")
)
