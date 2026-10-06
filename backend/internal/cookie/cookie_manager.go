package cookie

import (
	"context"
	"errors"
	"net/http"

	"github.com/1Vewton/MaterialScienceTV/backend/internal/ctxkey"
)

// GetResponseWriterFromContext gets response writer from context
func GetResponseWriterFromContext(
	ctx context.Context,
) (http.ResponseWriter, error) {
	rawWriter := ctx.Value(ctxkey.ResponseWriterKey)
	if rawWriter == nil {
		return nil, errors.New(
			"cannot get the response writer from the context",
		)
	}
	writer, ok := rawWriter.(http.ResponseWriter)
	if !ok {
		return nil, errors.New(
			"the value passed from context is not a response writer",
		)
	}
	return writer, nil
}

// GetRequestFromContext gets request from context
func GetRequestFromContext(
	ctx context.Context,
) (*http.Request, error) {
	rawRequest := ctx.Value(ctxkey.RequestKey)
	if rawRequest == nil {
		return nil, errors.New(
			"cannot get the reqeust from the context",
		)
	}
	request, ok := rawRequest.(*http.Request)
	if !ok {
		return nil, errors.New(
			"the value passed from context is not a request",
		)
	}
	return request, nil
}
