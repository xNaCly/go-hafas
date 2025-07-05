package gohafas

import gotime "time"

// Time wraps time.Time for some Hafas specific time and date interactions
type Time struct {
	inner gotime.Time
}

func TimeFrom(t gotime.Time) Time {
	return Time{t}
}

func (t *Time) ToHafasDateAndTime() (date string, time string) {
	date = t.inner.Format(gotime.DateOnly)
	time = t.inner.Format(gotime.TimeOnly)
	return
}
