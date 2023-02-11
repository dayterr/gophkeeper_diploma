package storage

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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
                                  card_number BYTEA NOT NULL, 
                                  exp_date BYTEA NOT NULL,
                                  cardholder BYTEA NOT NULL,
                                  cvv BYTEA NOT NULL,
                                  metadata text,
                                  user_id INT NOT NULL,
                                  FOREIGN KEY (user_id) REFERENCES public.users(id));`)
	if err != nil {
		return DBStorage{}, err
	}
	log.Info().Msg("table cards created")

	_, err = db.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS passwords (id serial PRIMARY KEY, 
                                  login BYTEA NOT NULL, 
                                  password BYTEA NOT NULL,
                                  metadata text,
                                  user_id INT NOT NULL,
                                  FOREIGN KEY (user_id) REFERENCES public.users(id));`)
	if err != nil {
		return DBStorage{}, err
	}
	log.Info().Msg("table passwords created")

	_, err = db.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS texts (id serial PRIMARY KEY, 
                                  text BYTEA NOT NULL,
                                  metadata text,
                                  user_id INT NOT NULL,
                                  FOREIGN KEY (user_id) REFERENCES public.users(id));`)
	if err != nil {
		return DBStorage{}, err
	}
	log.Info().Msg("table texts created")

	_, err = db.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS files (id serial PRIMARY KEY,
								  filename BYTEA NOT NULL,
                                  data BYTEA NOT NULL,
                                  metadata text,
                                  user_id INT NOT NULL,
                                  FOREIGN KEY (user_id) REFERENCES public.users(id));`)
	if err != nil {
		return DBStorage{}, err
	}
	log.Info().Msg("table files created")

	return DBStorage{
		DB:  db,
		DSN: dsn,
	}, nil
}
