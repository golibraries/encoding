package json

import (
	o "encoding/json"
)

type json struct{}

func (json) Marshal(v any) ([]byte, error) {
	return o.Marshal(v)
}

func (json) Unmarshal(data []byte, v any) error {
	return o.Unmarshal(data, v)
}

func New() *json {
	return &json{}
}

var e = New()

func Marshal(v any) ([]byte, error) {
	return e.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return e.Unmarshal(data, v)
}
