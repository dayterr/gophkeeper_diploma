package storage

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func (dbs DBStorage) AddText(ctx context.Context, userID int64, text Text) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Info().Msg("adding new text to database")

		_, err := dbs.DB.QueryContext(
			ctx,
			`INSERT INTO texts (text, metadata, user_id) VALUES ($1, $2, $3)`,
			text.Data, text.Metadata, userID)

		if err != nil {
			log.Info().Msg("error doing query")
			return err
		}
	}

	log.Info().Msg("text added successfully")

	return nil
}

func (dbs DBStorage) ListTexts(ctx context.Context, userID int64) ([]Text, error) {
	var texts []Text

	select {
	case <-ctx.Done():
		return []Text{}, ctx.Err()
	default:
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Info().Msg("getting texts from the database")
		res, err := dbs.DB.QueryContext(ctx,
			`SELECT id, text, metadata FROM texts WHERE user_id = $1`, userID)
		if err != nil {
			return []Text{}, err
		}

		defer res.Close()

		for res.Next() {
			text := Text{}
			err = res.Scan(&text.ID, &text.Data, &text.Metadata)
			if err != nil {
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

func (dbs DBStorage) DeleteText(ctx context.Context, userID, textID int64) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Info().Msg("deleting the text")
		log.Info().Msg(string(userID))
		log.Info().Msg(string(textID))
		res, err := dbs.DB.ExecContext(ctx,
			`DELETE FROM texts WHERE user_id = $1 and id = $2`, userID, textID)
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