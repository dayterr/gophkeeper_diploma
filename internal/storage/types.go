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
	CardNumber string
	ExpDate string
	Cardholder string
	CVV string
	Metadata string
}