package gohafas

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xNaCly/go-hafas/vbbraw"
)

func DerefIfNotNull[T any](t *T) T {
	if t != nil {
		return *t
	}
	var e T
	return e
}

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

func TestArrivalBoard(t *testing.T) {
	c, err := setup(t)
	assert.NoError(t, err)
	assert.NoError(t, c.Init())

	locations, err := c.LocationsByName("U Alexanderplatz", nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, locations)

	loc := locations[0]
	assert.NoError(t, loc.Unwrap())
	locAsStop, err := loc.AsStopLocation()
	assert.NoError(t, loc.Unwrap())

	arrivals, err := c.ArrivalBoard(*&locAsStop.Id, TimeFrom(time.Now()), nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, arrivals)
}
