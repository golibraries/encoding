package encoding

import (
	"errors"

	p "github.com/golang/protobuf/proto"
)

var (
	ErrProtoWrongValueType = errors.New("encoding proto converts on wrong type value")
)

type proto struct{}

func (proto) Marshal(v interface{}) ([]byte, error) {
	pb, ok := v.(p.Message)
	if !ok {
		return nil, ErrProtoWrongValueType
	}
	data, err := p.Marshal(pb)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (proto) Unmarshal(data []byte, v interface{}) error {
	pb, ok := v.(p.Message)
	if !ok {
		return ErrProtoWrongValueType
	}
	return p.Unmarshal(data, pb)
}
