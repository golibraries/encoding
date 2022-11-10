package protobuf

import (
	"errors"

	o "google.golang.org/protobuf/proto"
)

var (
	errProtobufWrongValueType = errors.New("protobuf: wrong type value")
)

type protobuf struct{}

func (protobuf) Marshal(v any) ([]byte, error) {
	pb, ok := v.(o.Message)
	if !ok {
		return nil, errProtobufWrongValueType
	}
	return o.Marshal(pb)
}

func (protobuf) Unmarshal(data []byte, v any) error {
	pb, ok := v.(o.Message)
	if !ok {
		return errProtobufWrongValueType
	}
	return o.Unmarshal(data, pb)
}

func New() *protobuf {
	return &protobuf{}
}

var e = New()

func Marshal(v any) ([]byte, error) {
	return e.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return e.Unmarshal(data, v)
}
