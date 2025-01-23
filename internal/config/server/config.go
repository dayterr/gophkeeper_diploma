package server

import (
	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func GetServerConfig() (ServerConfig, error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("getting config for server")

	var config ServerConfig
	err := env.Parse(&config)
	if err != nil {
		log.Info().Msg("error parsing env")
		return ServerConfig{}, err
	}

	if config.AddressCert == "" {
		return ServerConfig{}, ErrorNoAddressCert
	}

	if config.AddressCertKey == "" {
		return ServerConfig{}, ErrorNoAddressCertKey
	}

	if config.DatabaseDSN == "" {
		return ServerConfig{}, ErrorNoDSN
	}

	if config.JWTKey == "" {
		return ServerConfig{}, ErrorNoJWTKEY
	}

	log.Info().Msg("server config parsed")

	return config, nil
}
