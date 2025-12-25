package common

import "github.com/sunbay-developer/sunbay-nexus-sdk-go/constant"

// BaseResponse is the base response
type BaseResponse struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	TraceID string `json:"traceId,omitempty"`
}

// SetCode sets the response code
func (r *BaseResponse) SetCode(code string) {
	r.Code = code
}

// SetMsg sets the response message
func (r *BaseResponse) SetMsg(msg string) {
	r.Msg = msg
}

// SetTraceID sets the trace ID
func (r *BaseResponse) SetTraceID(traceID string) {
	r.TraceID = traceID
}

// IsSuccess checks if the response is successful
func (r *BaseResponse) IsSuccess() bool {
	return r.Code == constant.ResponseSuccessCode
}

