package vbbraw

import (
	"encoding/json"
	"slices"
	"strings"
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

// Must be called before calling LocationList_StopLocationOrCoordLocation_Item.Unwrap
func (l *LocationList_StopLocationOrCoordLocation_Item) Type() string {
	w := map[string]json.RawMessage{}
	if err := json.Unmarshal(l.union, &w); err != nil {
		return ""
	}
	for k := range w {
		return k
	}
	return ""
}

func (l TripType) AsGeoJSON() string {
	b := strings.Builder{}
	coords := make([][2]float64, 0, 16)
	for _, leg := range *l.LegList.Leg {
		if leg.GisRoute == nil || leg.GisRoute.Polyline == nil || leg.GisRoute.Polyline.Crd == nil {
			continue
		}
		for chunk := range slices.Chunk(*leg.GisRoute.Polyline.Crd, 2) {
			if len(chunk) != 2 {
				continue
			}
			coords = append(coords, [2]float64{chunk[0], chunk[1]})
		}
	}

	geo := map[string]any{
		"type": "Feature",
		"geometry": map[string]any{
			"type":        "LineString",
			"coordinates": coords,
		},
	}

	enc := json.NewEncoder(&b)
	enc.SetIndent("", "  ")
	_ = enc.Encode(geo)

	return b.String()
}
