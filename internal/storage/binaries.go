package storage

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func (dbs DBStorage) AddFile(ctx context.Context, login string, binary Binary) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Info().Msg("adding new file to database")

		_, err := dbs.DB.QueryContext(
			ctx,
			`INSERT INTO files (filename, data, metadata, user_id) VALUES ($1, $2, $3, (SELECT id FROM users WHERE login = $4))`,
			binary.Filename, binary.Data, binary.Metadata, login)

		if err != nil {
			log.Info().Msg("error doing query")
			return err
		}
	}

	log.Info().Msg("file added successfully")

	return nil
}

func (dbs DBStorage) ListFiles(ctx context.Context, login string) ([]Binary, error) {
	var binaries []Binary

	select {
	case <-ctx.Done():
		return []Binary{}, ctx.Err()
	default:
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Info().Msg("getting files from the database")
		res, err := dbs.DB.QueryContext(ctx,
			`SELECT id, filename, data, metadata FROM files WHERE user_id = (SELECT id FROM users WHERE login = $1)`, login)
		if err != nil {
			return []Binary{}, err
		}

		defer res.Close()

		for res.Next() {
			binary := Binary{}
			err = res.Scan(&binary.ID, &binary.Filename, &binary.Data, &binary.Metadata)
			if err != nil {
				return []Binary{}, err
			}

			binaries = append(binaries, binary)
		}

		if res.Err() != nil {
			return []Binary{}, err
		}
	}

	log.Info().Msg("files listed successfully")

	return binaries, nil
}

func (dbs DBStorage) DeleteFile(ctx context.Context, fileID int64, login string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Info().Msg("deleting the file")

		res, err := dbs.DB.ExecContext(ctx,
			`DELETE FROM files WHERE id = $1 AND user_id = (SELECT id FROM users WHERE login = $2)`, fileID, login)
		if err != nil {
			log.Info().Msg("error deleting file")
			return err
		}

		r, _ := res.RowsAffected()

		if r == 0 {
			return ErrorNotAuthorized
		}
	}

	log.Info().Msg("file deleted successfully")

	return nil
}
