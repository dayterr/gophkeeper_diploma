package storage

import (
	"context"
	"errors"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var ErrorNotAuthorized = errors.New("this user can't delete this data")

func (dbs DBStorage) AddCard(ctx context.Context, login string, card Card) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Info().Msg("adding new card to database")

		_, err := dbs.DB.QueryContext(
			ctx,
			`INSERT INTO cards (card_number, exp_date, cardholder, cvv, metadata, user_id) VALUES ($1, $2, $3, $4, $5, (SELECT id FROM users WHERE login = $6))`,
			card.CardNumber, card.ExpDate, card.Cardholder, card.CVV, card.Metadata, login)

		if err != nil {
			log.Info().Msg("error doing query")
			return err
		}
	}

	log.Info().Msg("card added successfully")
	return nil
}

func (dbs DBStorage) ListCards(ctx context.Context, login string) ([]Card, error) {
	var cards []Card

	select {
	case <-ctx.Done():
		return []Card{}, ctx.Err()
	default:
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Info().Msg("getting cards from the database")
		res, err := dbs.DB.QueryContext(ctx,
			`SELECT id, card_number, exp_date, cardholder, cvv, metadata FROM cards WHERE user_id = (SELECT id FROM users WHERE login = $1)`, login)
		if err != nil {
			return []Card{}, err
		}

		defer res.Close()

		for res.Next() {
			card := Card{}
			err = res.Scan(&card.ID, &card.CardNumber, &card.ExpDate, &card.Cardholder, &card.CVV, &card.Metadata)
			if err != nil {
				return []Card{}, err
			}

			cards = append(cards, card)
		}

		if res.Err() != nil {
			return []Card{}, err
		}
	}

	log.Info().Msg("cards listed successfully")
	return cards, nil
}

func (dbs DBStorage) DeleteCard(ctx context.Context, cardID int64, login string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Info().Msg("deleting the card")

		res, err := dbs.DB.ExecContext(ctx,
			`DELETE FROM cards WHERE id = $1 AND user_id = (SELECT id FROM users WHERE login = $2)`, cardID, login)
		if err != nil {
			log.Info().Msg("error deleting card")
			return err
		}

		r, _ := res.RowsAffected()

		if r == 0 {
			return ErrorNotAuthorized
		}
	}

	log.Info().Msg("card deleted successfully")

	return nil
}
