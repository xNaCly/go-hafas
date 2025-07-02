package gohafas

import "encoding/json"

type Error struct {
	ErrorCode string `json:"errorCode"`
	ErrorText string `json:"errorText"`
}

func (e *Error) Error() string {
	return "Hafas error: code=" + e.ErrorCode + ", text=" + e.ErrorText
}

func errorFromBytes(bytes []byte) *Error {
	var e Error
	json.Unmarshal(bytes, &e)
	return &e
}
