package vbbraw

import (
	"encoding/json"
)

// Unwrap extracts the value of the first element in a map
func Unwrap(wrapped json.RawMessage) (json.RawMessage, error) {
	w := map[string]json.RawMessage{}
	if err := json.Unmarshal(wrapped, &w); err != nil {
		return nil, err
	}
	for _, v := range w {
		return v, nil
	}
	return nil, nil
}

// Unwrap converts the wrapper { "TYPENAME": <real-content> } to <real-content>, by unmarshalling the wrapper
func (l *LocationList_StopLocationOrCoordLocation_Item) Unwrap() error {
	unwrapped, err := Unwrap(l.union)
	if err != nil {
		return err
	}
	l.union = unwrapped
	return nil
}
