package gohafas

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xNaCly/go-hafas/vbbraw"
)

func TestLocationsByName(t *testing.T) {
	c, err := setup(t)
	assert.NoError(t, err)
	assert.NoError(t, c.Init())

	maxNo := 1
	locations, err := c.LocationsByName("S Friedrichsstraße", &vbbraw.Verb8Params{
		MaxNo: &maxNo,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, locations)
}
