package main

import (
	"github.com/dayterr/gophkeeper_diploma/internal/handlers"
	"github.com/dayterr/gophkeeper_diploma/internal/routers"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/dayterr/gophkeeper_diploma/internal/config/server"
)


func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("starting server main")


	config, err := server.GetServerConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("please, set env vars")
	}

	log.Info().Msg("creating new handler")
	ah, err := handlers.NewAsyncHandler(config.DatabaseDSN, config.JWTKey, config.CryptoKey)
	if err != nil {
		log.Fatal().Err(err).Msg("error creating handler")
	}

	log.Info().Msg("creating new router")
	r := routers.CreateRouterWithAsyncHandler(ah)

	log.Info().Msg("starting secure server")
	err = http.ListenAndServeTLS(config.Address, config.AddressCert, config.AddressCertKey, r)
	if err != nil {
		log.Fatal().Err(err).Msg("server crashed")
	}
}