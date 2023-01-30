package storage

import (
	"context"
	"database/sql"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	_ "github.com/lib/pq"
	"time"
)

func NewDB(dsn string) (DBStorage, error) {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("creating database storage")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return DBStorage{}, err
	}

	_, err = db.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS users (id serial PRIMARY KEY, 
                                 login text UNIQUE NOT NULL, 
                                 password text NOT NULL);`)
	if err != nil {
		return DBStorage{}, err
	}
	log.Info().Msg("table users created")

	_, err = db.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS cards (id serial PRIMARY KEY, 
                                  card_number TEXT NOT NULL, 
                                  exp_date TEXT NOT NULL,
                                  cardholder TEXT NOT NULL,
                                  cvv TEXT NOT NULL,
                                  metadata text,
                                  user_id INT NOT NULL,
                                  FOREIGN KEY (user_id) REFERENCES public.users(id));`)
	if err != nil {
		return DBStorage{}, err
	}
	log.Info().Msg("table cards created")

	return DBStorage{
		DB:           db,
		DSN:          dsn,
	}, nil
}
