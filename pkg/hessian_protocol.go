package hessian2

import (
	"github.com/cloudwego/kitex/pkg/remote"
	"sync"
)

// must be strict read & strict write
var (
	bpPool sync.Pool
	_      BaseProtocol = (*BinaryProtocol)(nil)
)

func init() {
	bpPool.New = newBP
}

func newBP() interface{} {
	return &BinaryProtocol{}
}

// NewBinaryProtocol ...
func NewBinaryProtocol(t remote.ByteBuffer) *BinaryProtocol {
	bp := bpPool.Get().(*BinaryProtocol)
	bp.trans = t
	return bp
}

// BinaryProtocol ...
type BinaryProtocol struct {
	trans remote.ByteBuffer
	enc   Encoder
	dec   Decoder
}

// WriteByte ...
func (p *BinaryProtocol) WriteByte(value int8) error {
	err := p.enc.Encode(value)
	return err
}

// ReadByte ...
func (p *BinaryProtocol) ReadByte() (value int8, err error) {
	err = p.dec.Decode(p.trans)
	if err != nil {
		return 0, err
	}
	readByte, _ := p.dec.buffer.ReadByte()
	return int8(readByte), err
}
