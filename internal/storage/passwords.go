package storage

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func (dbs DBStorage) AddPassword(ctx context.Context, login string, password Password) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Info().Msg("adding new password to database")

		_, err := dbs.DB.QueryContext(
			ctx,
			`INSERT INTO passwords (login, password, metadata, user_id) VALUES ($1, $2, $3, (SELECT id FROM users WHERE login = $4))`,
			password.Login, password.Password, password.Metadata, login)

		if err != nil {
			log.Info().Msg("error doing query")
			return err
		}
	}

	log.Info().Msg("password added successfully")

	return nil
}

func (dbs DBStorage) ListPasswords(ctx context.Context, login string) ([]Password, error) {
	var passwords []Password

	select {
	case <-ctx.Done():
		return []Password{}, ctx.Err()
	default:
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Info().Msg("getting passwords from the database")
		res, err := dbs.DB.QueryContext(ctx,
			`SELECT id, login, password, metadata FROM passwords WHERE user_id = (SELECT id FROM users WHERE login = $1)`, login)
		if err != nil {
			return []Password{}, err
		}

		defer res.Close()

		for res.Next() {
			password := Password{}
			err = res.Scan(&password.ID, &password.Login, &password.Password, &password.Metadata)
			if err != nil {
				return []Password{}, err
			}

			passwords = append(passwords, password)
		}

		if res.Err() != nil {
			return []Password{}, err
		}
	}

	log.Info().Msg("passwords listed successfully")
	return passwords, nil
}

func (dbs DBStorage) DeletePassword(ctx context.Context, passwordID int64, login string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Info().Msg("deleting the card")

		res, err := dbs.DB.ExecContext(ctx,
			`DELETE FROM passwords WHERE id = $1 AND user_id = (SELECT id FROM users WHERE login = $2)`, passwordID, login)
		if err != nil {
			log.Info().Msg("error deleting password")
			return err
		}

		r, _ := res.RowsAffected()

		if r == 0 {
			return ErrorNotAuthorized
		}
	}

	log.Info().Msg("password deleted successfully")

	return nil
}
