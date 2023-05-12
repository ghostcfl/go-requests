package requests

import "errors"

const (
	defaultUserAgent = "go-requests/1.0.3"

	GET     = "GET"
	POST    = "POST"
	PUT     = "PUT"
	DELETE  = "DELETE"
	OPTIONS = "OPTIONS"
	HEAD    = "HEAD"
	PATCH   = "PATCH"
)

var (
	// ErrHookFuncMaxLimit will be throwed when the number of hook functions
	// more than MaxLimit = 8
	ErrHookFuncMaxLimit = errors.New("HOOK函数数量要少于8个")

	// ErrIndexOutofBound means the index out of bound
	ErrIndexOutOfBound = errors.New("Index超出")
)
