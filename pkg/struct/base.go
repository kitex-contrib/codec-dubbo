package base

type Base struct {
}

func NewBaseStruct() *Base {
	return &Base{}
}

func (b Base) Read() error {
	return nil
}

func (b Base) Write() error {
	return nil
}
