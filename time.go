package gohafas

import "time"

// Time wraps time.Time for some Hafas specific time and date interactions
type Time struct {
	inner time.Time
}

func TimeFrom(t time.Time) Time {
	return Time{t}
}

func (t Time) ToHafasDateAndTime() (dateStr string, timeStr string) {
	dateStr = t.inner.Format(time.DateOnly)
	timeStr = t.inner.Format(time.TimeOnly)
	return
}
