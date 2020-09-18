package rpc

import (
	"google.golang.org/grpc/codes"
)

const (
	InnerCodeOK      = 0
	InnerCodeUnknown = 99999
)

type Error struct {
	innerCode int
	rpcCode   codes.Code
	text      string
}

func NewError(innerCode int, rpcCode codes.Code, text string) *Error {
	return &Error{
		innerCode: innerCode,
		rpcCode:   rpcCode,
		text:      text,
	}
}

func (e Error) InnerCode() int {
	return e.innerCode
}

func (e Error) RPCCode() codes.Code {
	return e.rpcCode
}

func (e Error) Error() string {
	return e.text
}

func innerCode(err error) int {
	if err == nil {
		return InnerCodeOK
	}

	code := InnerCodeUnknown
	if err, ok := err.(interface {
		InnerCode() int
	}); ok {
		code = err.InnerCode()
	}

	return code
}

func rpcCode(err error) codes.Code {
	if err == nil {
		return codes.OK
	}

	grpcCode := codes.Internal
	if err, ok := err.(interface {
		RPCCode() codes.Code
	}); ok {
		grpcCode = err.RPCCode()
	}

	return grpcCode
}
