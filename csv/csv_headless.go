package csv

import (
	"bytes"

	o "github.com/gocarina/gocsv"
)

type csvheadless struct{}

func (csvheadless) Marshal(v interface{}) ([]byte, error) {
	var buf = new(bytes.Buffer)
	err := o.MarshalWithoutHeaders(v, buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (csvheadless) Unmarshal(data []byte, v interface{}) error {
	return o.UnmarshalWithoutHeaders(bytes.NewBuffer(data), v)
}

func NewHeadless() *csvheadless {
	return &csvheadless{}
}

var eHeadless = NewHeadless()

func MarshalHeadless(v interface{}) ([]byte, error) {
	return e.Marshal(v)
}

func UnmarshalHeadless(data []byte, v interface{}) error {
	return e.Unmarshal(data, v)
}
