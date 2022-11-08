package csv

import (
	"bytes"

	o "github.com/gocarina/gocsv"
)

type csv struct{}

func (csv) Marshal(v interface{}) ([]byte, error) {
	var buf = new(bytes.Buffer)
	err := o.Marshal(v, buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (csv) Unmarshal(data []byte, v interface{}) error {
	return o.Unmarshal(bytes.NewBuffer(data), v)
}

func New() *csv {
	return &csv{}
}

var e = New()

func Marshal(v interface{}) ([]byte, error) {
	return e.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return e.Unmarshal(data, v)
}
