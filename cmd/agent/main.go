package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/dayterr/gophkeeper_diploma/internal/config/agent"
	"github.com/dayterr/gophkeeper_diploma/internal/tuiagent"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("starting client main")

	log.Info().Msg("getting agent config in the main")
	config, err := agent.GetAgentConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("please, set env vars")
	}

	tuiClient, err := tuiagent.NewTUICLient(config.AddressCert, config.AddressCertKey, config.Address)
	if err != nil {
		log.Fatal().Err(err).Msg("creating tui client failed")
	}

	tuiClient.Run()
}
