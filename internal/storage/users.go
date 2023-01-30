package storage

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func (dbs DBStorage) AddUser(ctx context.Context, user User) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Info().Msg("adding new user to database")

		_, err := dbs.DB.QueryContext(ctx,
			`INSERT INTO users (login, password) VALUES ($1, $2)`,
			user.Login, user.Password)

		if err != nil {
			return err
		}
	}

	log.Info().Msg("user added successfully")
	return nil
}

func (dbs DBStorage) GetUser(ctx context.Context, login string) (int64, error) {
	var userID int64

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Info().Msg("getting user from database")

		res := dbs.DB.QueryRowContext(ctx, `SELECT id FROM users WHERE login = $1`, login)
		err := res.Scan(&userID)
		if err != nil {
			return 0, err
		}
	}

	log.Info().Msg("user found successfully")

	return userID, nil
}
