package xml

import (
	o "encoding/xml"
)

type xml struct{}

func (xml) Marshal(v any) ([]byte, error) {
	return o.Marshal(v)
}

func (xml) Unmarshal(data []byte, v any) error {
	return o.Unmarshal(data, v)
}

func New() *xml {
	return &xml{}
}

var e = New()

func Marshal(v any) ([]byte, error) {
	return e.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return e.Unmarshal(data, v)
}
