package yaml

import (
	o "gopkg.in/yaml.v3"
)

type yaml struct{}

func (yaml) Marshal(v any) ([]byte, error) {
	return o.Marshal(v)
}

func (yaml) Unmarshal(data []byte, v any) error {
	return o.Unmarshal(data, v)
}

func New() *yaml {
	return &yaml{}
}

var e = New()

func Marshal(v any) ([]byte, error) {
	return e.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return e.Unmarshal(data, v)
}
