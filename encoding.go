package encoding

type (

	// Marshaler represents a marshal interface
	Marshaler interface {
		Marshal(any) ([]byte, error)
	}

	// Unmarshaler represents a Unmarshal interface
	Unmarshaler interface {
		Unmarshal([]byte, any) error
	}

	// Encoding is the interface that groups the basic Marshal and Unmarshal methods.
	Encoding interface {
		Marshaler
		Unmarshaler
	}
)

type Identifier int

func (e Identifier) Marshal(v any) ([]byte, error) {
	return encodings[int(e)].Marshal(v)
}

func (e Identifier) Unmarshal(data []byte, v any) error {
	return encodings[int(e)].Unmarshal(data, v)
}

const ()

var encodings = [...]Encoding{}
