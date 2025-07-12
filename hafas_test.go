package gohafas

import (
	"context"
	"log/slog"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xnacly/go-hafas/language"
	"github.com/xnacly/go-hafas/vbbraw"
)

func TestReadme(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	httpClient := http.Client{
		Timeout: 2 * time.Second,
	}

	client, err := NewClient(
		getStringFromEnvOrFailTest(t, BASEURL_ENV),
		getStringFromEnvOrFailTest(t, AUTH_ENV),
		WithLanguage(language.ES),
		WithContext(ctx),
		WithHttpClient(&httpClient),
	)

	if err != nil {
		slog.Error("hafas", "msg", "failed to create hafas client", "err", err)
	}

	err = client.Init()
	if err != nil {
		slog.Error("hafas", "msg", "failed to init hafas client", "err", err)
	}

	err = client.Ping()
	if err != nil {
		slog.Error("hafas", "msg", "failed to ping hafas remote", "err", err)
	}
}

func derefIfNotNull[T any](t *T) T {
	if t != nil {
		return *t
	}
	panic("Cant deref, argument was nil")
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

func TestLocationsByCoordinates(t *testing.T) {
	c, err := setup(t)
	assert.NoError(t, err)
	assert.NoError(t, c.Init())

	maxNo := 1
	locations, err := c.LocationsByName("S Friedrichsstraße", &vbbraw.Verb8Params{
		MaxNo: &maxNo,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, locations)

	loc := locations[0]
	assert.NoError(t, loc.Unwrap())
	coord, err := loc.AsStopLocation()
	assert.NoError(t, err)

	stopLocation, err := c.LocationsByCoordinate(derefIfNotNull(coord.Lat), derefIfNotNull(coord.Lon), nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, stopLocation)
}

func TestArrivals(t *testing.T) {
	c, err := setup(t)
	assert.NoError(t, err)
	assert.NoError(t, c.Init())

	locations, err := c.LocationsByName("U Alexanderplatz", nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, locations)

	loc := locations[0]
	assert.NoError(t, loc.Unwrap())
	locAsStop, err := loc.AsStopLocation()

	arrivals, err := c.Arrivals(locAsStop.Id, TimeFrom(time.Now()), nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, arrivals)
}

func TestDepartures(t *testing.T) {
	c, err := setup(t)
	assert.NoError(t, err)
	assert.NoError(t, c.Init())

	locations, err := c.LocationsByName("U Alexanderplatz", nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, locations)

	loc := locations[0]
	assert.NoError(t, loc.Unwrap())
	locAsStop, err := loc.AsStopLocation()
	departure, err := c.Departures(locAsStop.Id, TimeFrom(time.Now()), nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, departure)
}

func TestDataInfo(t *testing.T) {
	c, err := setup(t)
	assert.NoError(t, err)
	assert.NoError(t, c.Init())

	data, err := c.DataInfo()
	assert.NotEmpty(t, data)
}

func TestTrip(t *testing.T) {
	c, err := setup(t)
	assert.NoError(t, err)
	assert.NoError(t, c.Init())

	locations, err := c.LocationsByName("U Alexanderplatz", nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, locations)
	assert.NoError(t, locations[0].Unwrap())
	start, err := locations[0].AsStopLocation()
	assert.NoError(t, err)

	locations, err = c.LocationsByName("U Jannowitzbrücke", nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, locations)
	assert.NoError(t, locations[0].Unwrap())
	end, err := locations[0].AsStopLocation()
	assert.NoError(t, err)

	params := &vbbraw.Verb11Params{
		OriginId: &start.Id,
		DestId:   &end.Id,
	}
	trip, err := c.TripSearch(TimeFrom(time.Now()), params)
	assert.NoError(t, err)
	assert.NotEmpty(t, trip)
}

func TestJourneyDetail(t *testing.T) {
	c, err := setup(t)
	assert.NoError(t, err)
	assert.NoError(t, c.Init())

	locations, err := c.LocationsByName("S Köpenick", nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, locations)

	loc := locations[0]
	assert.NoError(t, loc.Unwrap())
	locAsStop, err := loc.AsStopLocation()

	arrivals, err := c.Arrivals(locAsStop.Id, TimeFrom(time.Now()), nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, arrivals)

	journey, err := c.JourneyDetail(arrivals[0].JourneyDetailRef.Ref, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, journey)
}

func TestJourneyPos(t *testing.T) {
	c, err := setup(t)
	assert.NoError(t, err)
	assert.NoError(t, c.Init())

	locations, err := c.LocationsByName("U Adenauerplatz", nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, locations)

	loc := locations[0]
	assert.NoError(t, loc.Unwrap())
	locAsStop, err := loc.AsStopLocation()

	lat, lon := derefIfNotNull(locAsStop.Lat), derefIfNotNull(locAsStop.Lon)
	// ±0.01 is equal to around 1.1km, so we define a bounding box  of around 2km*2km around the
	// stop we looked up above
	offset := float32(0.01)
	journeys, err := c.JourneyPos(lat-offset, lon-offset, lat+offset, lon+offset, TimeFrom(time.Now()), nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, journeys)
}

func TestHimSearch(t *testing.T) {
	c, err := setup(t)
	assert.NoError(t, err)
	assert.NoError(t, c.Init())

	// start 6 months ago
	dateB, _ := TimeFrom(time.Now().AddDate(0, -6, 0)).ToHafasDateAndTime()
	dateE, _ := TimeFrom(time.Now()).ToHafasDateAndTime()

	param := &vbbraw.Verb5Params{
		DateB: &dateB,
		DateE: &dateE,
	}
	hims, err := c.HimSearch(param)
	assert.NoError(t, err)

	// TODO: i have no idea if this works, when testing there were no him
	// messages available, maybe the test env has none?
	debugStruct(hims)
}
