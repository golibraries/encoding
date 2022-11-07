package encoding

import (
	originJSON "encoding/json"
)

type json struct{}

func (json) Marshal(v interface{}) ([]byte, error) {
	return originJSON.Marshal(v)
}

func (json) Unmarshal(data []byte, v interface{}) error {
	return originJSON.Unmarshal(data, v)
}
