package rpc

import (
	"google.golang.org/grpc/codes"
	"net/http"
)

func HTTPStatus(err error) int {
	if err == nil {
		return http.StatusOK
	}

	grpcCode := rpcCode(err)
	return mapHTTPCode(grpcCode)
}

func HTTPErrCode(err error) int {
	if err == nil {
		return InnerCodeUnknown
	}

	return innerCode(err)
}

func mapHTTPCode(grpcCode codes.Code) int {
	switch grpcCode {
	case codes.OK:
		return http.StatusOK
	case codes.FailedPrecondition, codes.InvalidArgument, codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.NotFound:
		return http.StatusNotFound
	case codes.Aborted, codes.AlreadyExists:
		return http.StatusConflict
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.Canceled:
		return 499
	case codes.Unknown, codes.Internal, codes.DataLoss:
		return http.StatusInternalServerError
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	default:
		return http.StatusInternalServerError
	}
}