package commons

type ProtocolType string

const (
	TTHeader ProtocolType = "TTHeader"
	HTTP2    ProtocolType = "HTTP2"
)

func (p ProtocolType) String() string {
	switch p {
	case TTHeader:
		return "TTHeader"
	case HTTP2:
		return "HTTP2"
	default:
		return ""
	}
}

func (p ProtocolType) Set(s string) error {
	panic("this operation is forbidden")
}
