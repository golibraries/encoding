package json

import (
	o "encoding/json"
)

type jsonindent struct{}

func (jsonindent) Marshal(v any) ([]byte, error) {
	return o.MarshalIndent(v, "", "  ")
}

func (jsonindent) Unmarshal(data []byte, v any) error {
	return o.Unmarshal(data, v)
}

func NewIndent() *jsonindent {
	return &jsonindent{}
}

var eIndent = NewIndent()

func MarshalIndent(v any) ([]byte, error) {
	return eIndent.Marshal(v)
}

func UnmarshalIndent(data []byte, v any) error {
	return eIndent.Unmarshal(data, v)
}
