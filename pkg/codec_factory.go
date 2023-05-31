package hessian2

import (
	"github.com/kitex-contrib/codec-hessian2/pkg/codec"
	commons "github.com/kitex-contrib/codec-hessian2/pkg/common"
)

func getCodec(protocolType commons.ProtocolType) (c Codec) {
	switch protocolType {
	case commons.TTHeader:
	case commons.HTTP2:
	default:
		c = codec.NewDefaultCodec()
	}
	return c
}
