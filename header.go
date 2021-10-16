package clickup

type header string

func (h header) String() string { return string(h) }

const (
	_authorization header = "Authorization"
	_contentType   header = "Content-Type"
	_xSignature    header = "X-Signature"
)
