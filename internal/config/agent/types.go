package agent

type AgentConfig struct {
	Address        string `env:"ADDRESS" envDefault:"localhost:8080"`
	AddressCert    string `env:"ADDRESS_CERT"`
	AddressCertKey string `env:"ADDRESS_CERT_KEY"`
}
