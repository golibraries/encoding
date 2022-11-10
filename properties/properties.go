package properties

import (
	o "github.com/magiconair/properties"
)

type properties struct{}

func (properties) Marshal(v any) ([]byte, error) {
	// TODO: implement
	return nil, nil
}

func (properties) Unmarshal(data []byte, v any) error {
	p, err := o.LoadString(string(data))
	if err != nil {
		return err
	}
	return p.Decode(v)
}

func New() *properties {
	return &properties{}
}

var e = New()

func Marshal(v any) ([]byte, error) {
	return e.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return e.Unmarshal(data, v)
}
