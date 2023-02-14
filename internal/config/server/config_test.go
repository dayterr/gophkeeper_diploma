package server

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetServerConfig(t *testing.T) {
	testIP := "89.76.173.44:8001"
	testCert := "somefolder/ca.cert.pem"
	testCertKey := "somefolder/ca.key.pem"
	testDSN := "postgres://postgres:somepass@localhost:5432/pdb?sslmode=disable"
	testJWTKey := "nfejf4556.dsff"

	os.Setenv("ADDRESS", testIP)
	os.Setenv("DATABASE_DSN", testDSN)
	os.Setenv("JWT_KEY", testJWTKey)
	os.Setenv("ADDRESS_CERT", testCert)
	os.Setenv("ADDRESS_CERT_KEY", testCertKey)

	resConfig, err := GetServerConfig()
	assert.NoError(t, err)
	assert.Equal(t, resConfig.Address, testIP, "Problems parsing env")
	assert.Equal(t, resConfig.DatabaseDSN, testDSN, "Problems parsing env")
	assert.Equal(t, resConfig.JWTKey, testJWTKey, "Problems parsing env")
	assert.Equal(t, resConfig.AddressCert, testCert, "Problems parsing env")
	assert.Equal(t, resConfig.AddressCertKey, testCertKey, "Problems parsing env")

	os.Unsetenv("DATABASE_DSN")
	_, err = GetServerConfig()
	assert.ErrorIs(t, err, ErrorNoDSN)

	os.Setenv("DATABASE_DSN", testDSN)
	os.Unsetenv("JWT_KEY")
	_, err = GetServerConfig()
	assert.ErrorIs(t, err, ErrorNoJWTKEY)

	os.Setenv("JWT_KEY", testJWTKey)
	os.Unsetenv("ADDRESS_CERT")
	_, err = GetServerConfig()
	assert.ErrorIs(t, err, ErrorNoAddressCert)

	os.Setenv("ADDRESS_CERT", testCert)
	os.Unsetenv("ADDRESS_CERT_KEY")
	_, err = GetServerConfig()
	assert.ErrorIs(t, err, ErrorNoAddressCertKey)
}
