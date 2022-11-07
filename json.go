package encoding

import (
	j "encoding/json"
)

type json struct{}

func (json) Marshal(v interface{}) ([]byte, error) {
	return j.Marshal(v)
}

func (json) Unmarshal(data []byte, v interface{}) error {
	return j.Unmarshal(data, v)
}
