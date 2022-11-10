package fury

type fury struct{}

func (fury) Marshal(v interface{}) ([]byte, error) {
	// TODO: implement
	return nil, nil
}

func (fury) Unmarshal(data []byte, v interface{}) error {
	// TODO: implement
	return nil
}

func New() *fury {
	return &fury{}
}

var e = New()

func Marshal(v interface{}) ([]byte, error) {
	return e.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return e.Unmarshal(data, v)
}
