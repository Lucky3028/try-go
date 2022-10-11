package app_errors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Lucky3028/try-go/api/middlewares"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err error) {
	var appErr *ApplicationError
	if !errors.As(err, &appErr) {
		appErr = &ApplicationError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	traceId := middlewares.GetTraceId(request.Context())
	log.Printf("[%d]error: %s\n", traceId, appErr)

	var statusCode int

	switch appErr.ErrCode {
	case DataNotFound:
		statusCode = http.StatusNotFound
	case NoTargetData, RequestBodyDecodeFailed, BadParam:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	writer.WriteHeader(statusCode)
	json.NewEncoder(writer).Encode(appErr)
}
