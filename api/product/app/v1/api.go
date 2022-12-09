package v1

import (
	"context"
	"net/http"
)

type HTTPResponse struct {
	Code    int
	Message string
	Data    interface{}
}

type TextHTTPServer interface {
	QueryTextByFilter(ctx context.Context, reqeust *http.Request) (*HTTPResponse, error)
}
