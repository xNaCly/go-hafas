package gohafas

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	BASEURL_ENV = "BASEURL"
	AUTH_ENV    = "AUTH"
)

func getStringFromEnvOrFailTest(t *testing.T, name string) string {
	envVal := os.Getenv(name)
	assert.NotEmpty(t, envVal)
	return envVal
}

func setup(t *testing.T) (*Client, error) {
	t.Helper()
	baseUrl := getStringFromEnvOrFailTest(t, BASEURL_ENV)
	auth := getStringFromEnvOrFailTest(t, AUTH_ENV)
	return NewClient(baseUrl, auth)
}

func TestNewClient(t *testing.T) {
	c, err := setup(t)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestNewClientInit(t *testing.T) {
	c, err := setup(t)
	assert.NoError(t, err)
	assert.NotNil(t, c)
	err = c.Init()
	assert.NoError(t, err)
}

func TestNewClientInitAndPing(t *testing.T) {
	c, err := setup(t)
	assert.NoError(t, err)
	assert.NotNil(t, c)
	err = c.Init()
	assert.NoError(t, err)
	err = c.Ping()
	assert.NoError(t, err)
}
