package flatbuffers

import (
	o "github.com/google/flatbuffers/go"
)

type flatbuffers struct{}

var _codec = &o.FlatbuffersCodec{}

func (flatbuffers) Marshal(v any) ([]byte, error) {
	return _codec.Marshal(v)
}

func (flatbuffers) Unmarshal(data []byte, v any) error {
	return _codec.Unmarshal(data, v)
}

func New() *flatbuffers {
	return &flatbuffers{}
}

var e = New()

func Marshal(v any) ([]byte, error) {
	return e.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return e.Unmarshal(data, v)
}
