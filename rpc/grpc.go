package rpc

import (
	protobuf "github.com/light-service/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GRPCStatus(err error) codes.Code {
	return status.Code(err)
}

func GRPCErrCode(err error) int {
	if err == nil {
		return InnerCodeUnknown
	}

	return innerCode(err)
}

func GRPCError(err error) error {
	if err == nil {
		return nil
	}
	rpcCode := rpcCode(err)
	if rpcCode == codes.OK {
		return nil
	}
	innerCode := innerCode(err)

	s := status.New(rpcCode, err.Error())
	if innerCode != InnerCodeOK {
		sd, err := s.WithDetails(&protobuf.Error{
			Code: int32(innerCode),
			Msg:  err.Error(),
		})
		if err == nil {
			s = sd
		}
	}
	return s.Err()
}
