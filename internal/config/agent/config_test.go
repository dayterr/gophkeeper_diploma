package agent

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAgentConfig(t *testing.T) {
	testIP := "89.76.173.44:8001"
	testCert := "somefolder/ca.cert.pem"
	testCertKey := "somefolder/ca.key.pem"

	os.Setenv("ADDRESS", testIP)
	os.Setenv("ADDRESS_CERT", testCert)
	os.Setenv("ADDRESS_CERT_KEY", testCertKey)

	resConfig, err := GetAgentConfig()
	assert.NoError(t, err)
	assert.Equal(t, resConfig.Address, testIP, "Problems parsing env")
	assert.Equal(t, resConfig.AddressCert, testCert, "Problems parsing env")
	assert.Equal(t, resConfig.AddressCertKey, testCertKey, "Problems parsing env")

	os.Unsetenv("ADDRESS")
	resConfig, err = GetAgentConfig()
	assert.NoError(t, err)
	assert.Equal(t, resConfig.Address, DefaultAddress, "IP should be equal to the default value")

	os.Unsetenv("ADDRESS_CERT")
	_, err = GetAgentConfig()
	assert.Error(t, err)

	os.Setenv("ADDRESS_CERT", testCert)
	os.Unsetenv("ADDRESS_CERT_KEY")
	_, err = GetAgentConfig()
	assert.Error(t, err)

}