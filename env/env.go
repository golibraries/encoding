package env

type env struct{} // auto sort keys

func (env) Marshal(v any) ([]byte, error) {
	// TODO: implement
	return nil, nil
}

func (env) Unmarshal(data []byte, v any) error {
	// TODO: implement
	return nil
}

func New() *env {
	return &env{}
}

var e = New()

func Marshal(v any) ([]byte, error) {
	return e.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return e.Unmarshal(data, v)
}
