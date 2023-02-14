package storage

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func (dbs DBStorage) AddText(ctx context.Context, login string, text Text) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Info().Msg("adding new text to database")

		_, err := dbs.DB.QueryContext(
			ctx,
			`INSERT INTO texts (text, metadata, user_id) VALUES ($1, $2, (SELECT id FROM users WHERE login = $3))`,
			text.Data, text.Metadata, login)

		if err != nil {
			log.Info().Msg("error doing query")
			return err
		}
	}

	log.Info().Msg("text added successfully")

	return nil
}

func (dbs DBStorage) ListTexts(ctx context.Context, login string) ([]Text, error) {
	var texts []Text

	select {
	case <-ctx.Done():
		return []Text{}, ctx.Err()
	default:
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Info().Msg("getting texts from the database")
		res, err := dbs.DB.QueryContext(ctx,
			`SELECT id, text, metadata FROM texts WHERE user_id = (SELECT id FROM users WHERE login = $1)`, login)
		if err != nil {
			log.Info().Msg(err.Error())
			return []Text{}, err
		}

		defer res.Close()

		for res.Next() {
			text := Text{}
			err = res.Scan(&text.ID, &text.Data, &text.Metadata)
			if err != nil {
				log.Info().Msg(err.Error())
				return []Text{}, err
			}

			texts = append(texts, text)
		}

		if res.Err() != nil {
			return []Text{}, err
		}
	}

	log.Info().Msg("texts listed successfully")

	return texts, nil
}

func (dbs DBStorage) DeleteText(ctx context.Context, textID int64, login string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Info().Msg("deleting the text")
		res, err := dbs.DB.ExecContext(ctx,
			`DELETE FROM texts WHERE id = $1 AND user_id = (SELECT id FROM users WHERE login = $2)`, textID, login)
		if err != nil {
			log.Info().Msg("error deleting text")
			return err
		}

		r, _ := res.RowsAffected()

		if r == 0 {
			return ErrorNotAuthorized
		}
	}

	log.Info().Msg("text deleted successfully")

	return nil
}
