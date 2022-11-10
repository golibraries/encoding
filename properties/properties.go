package properties

import (
	"bytes"
	"sort"

	o "github.com/magiconair/properties"

	"github.com/golibraries/cast"
)

type properties struct{}

func (properties) Marshal(v any) ([]byte, error) {
	m, err := cast.ToFlatStringMapE(v)
	if err != nil {
		return nil, err
	}

	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	p := o.NewProperties()
	for _, key := range keys {
		_, _, err := p.Set(key, cast.ToString(m[key]))
		if err != nil {
			return nil, err
		}
	}

	var buf bytes.Buffer

	_, err = p.WriteComment(&buf, "#", o.UTF8)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
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
