package hessian2

import "context"

type Message struct {
}

type ByteBuffer struct {
}

// Codec is the abstraction of the codec layer of Kitex.
type Codec interface {
	Encode(ctx context.Context, msg Message, out ByteBuffer) error

	Decode(ctx context.Context, msg Message, in ByteBuffer) error

	Name() string
}
