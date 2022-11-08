package csvheadless

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

func New() *csvheadless {
	return &csvheadless{}
}

var e = New()

func Marshal(v interface{}) ([]byte, error) {
	return e.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return e.Unmarshal(data, v)
}
