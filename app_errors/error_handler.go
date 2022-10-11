package app_errors

import (
	"encoding/json"
	"errors"
	"net/http"
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
