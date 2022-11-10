package ini

type ini struct{}

func (ini) Marshal(v interface{}) ([]byte, error) {
	// TODO: implement
	return nil, nil
}

func (ini) Unmarshal(data []byte, v interface{}) error {
	// TODO: implement
	return nil
}

func New() *ini {
	return &ini{}
}

var e = New()

func Marshal(v interface{}) ([]byte, error) {
	return e.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return e.Unmarshal(data, v)
}
