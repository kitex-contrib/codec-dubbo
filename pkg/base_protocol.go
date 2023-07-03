package hessian2

type BaseProtocol interface {
	// WriteByte ...
	WriteByte(value int8) error

	// ReadByte ...
	ReadByte() (value int8, err error)
}
