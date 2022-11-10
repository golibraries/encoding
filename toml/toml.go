package toml

import (
	o "github.com/pelletier/go-toml/v2"
)

type toml struct{}

func (toml) Marshal(v interface{}) ([]byte, error) {
	return o.Marshal(v)
}

func (toml) Unmarshal(data []byte, v interface{}) error {
	return o.Unmarshal(data, v)
}

func New() *toml {
	return &toml{}
}

var e = New()

func Marshal(v interface{}) ([]byte, error) {
	return e.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return e.Unmarshal(data, v)
}
