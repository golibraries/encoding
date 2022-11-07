package encoding

type (

	// Marshaler represents a marshal interface
	Marshaler interface {
		Marshal(interface{}) ([]byte, error)
	}

	// Unmarshaler represents a Unmarshal interface
	Unmarshaler interface {
		Unmarshal([]byte, interface{}) error
	}

	// Encoding is the interface that groups the basic Marshal and Unmarshal methods.
	Encoding interface {
		Marshaler
		Unmarshaler
	}
)

type Identifier int

func (e Identifier) Marshal(v interface{}) ([]byte, error) {
	return encodings[int(e)].Marshal(v)
}

func (e Identifier) Unmarshal(data []byte, v interface{}) error {
	return encodings[int(e)].Unmarshal(data, v)
}

func (e Identifier) Encoding() Encoding {
	return encodings[int(e)]
}

const (
	JSON Identifier = iota
)

var encodings = [...]Encoding{
	JSON: json{},
}
