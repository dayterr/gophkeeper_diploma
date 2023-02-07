package server

type ServerConfig struct {
	Address string `env:"ADDRESS" envDefault:"localhost:8080"`
	DatabaseDSN   string        `env:"DATABASE_DSN"`
	JWTKey string `env:"JWT_KEY"`
	AddressCert string `env:"ADDRESS_CERT"`
	AddressCertKey string `env:"ADDRESS_CERT_KEY"`
	CryptoKey string `env:"CRYPTO_KEY"`
}
