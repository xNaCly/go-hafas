package gohafas

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getStringFromEnvOrFailTest(t *testing.T, name string) string {
	envVal := os.Getenv(name)
	assert.NotEmpty(t, envVal)
	return envVal
}

const (
	BASEURL_ENV = "BASEURL"
	AUTH_ENV    = "AUTH"
)

func TestNewClient(t *testing.T) {
	baseUrl := getStringFromEnvOrFailTest(t, BASEURL_ENV)
	auth := getStringFromEnvOrFailTest(t, AUTH_ENV)

	c, err := NewClient(baseUrl, auth)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}
