package hcl

type hcl struct{}

func (hcl) Marshal(v any) ([]byte, error) {
	// TODO: implement
	return nil, nil
}

func (hcl) Unmarshal(data []byte, v any) error {
	// TODO: implement
	return nil
}

func New() *hcl {
	return &hcl{}
}

var e = New()

func Marshal(v any) ([]byte, error) {
	return e.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return e.Unmarshal(data, v)
}
