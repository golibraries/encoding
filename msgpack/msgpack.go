package msgpack

import (
	o "github.com/vmihailenco/msgpack/v5"
)

type msgpack struct{}

func (msgpack) Marshal(v any) ([]byte, error) {
	return o.Marshal(v)
}

func (msgpack) Unmarshal(data []byte, v any) error {
	return o.Unmarshal(data, v)
}

func New() *msgpack {
	return &msgpack{}
}

var e = New()

func Marshal(v any) ([]byte, error) {
	return e.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return e.Unmarshal(data, v)
}
