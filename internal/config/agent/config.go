package agent

import (
	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const DefaultAddress = "localhost:443"

func GetAgentConfig() (AgentConfig, error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("getting config for agent")

	var config AgentConfig
	err := env.Parse(&config)
	if err != nil {
		log.Info().Msg("error parsing env")
		return AgentConfig{}, err
	}

	if config.AddressCert == "" {
		return AgentConfig{}, ErrorNoAddressCert
	}

	if config.AddressCertKey == "" {
		return AgentConfig{}, ErrorNoAddressCertKey
	}

	log.Info().Msg("agent config parsed")

	return config, nil
}
