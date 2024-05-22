package utils

import "github.com/zakirkun/test-elabram/internal/globals"

func SetGeneralResponse(status, message, responseId string, data any) globals.GeneralResponse {
	return globals.GeneralResponse{
		Status:     status,
		Message:    message,
		ResponseId: responseId,
		Data:       data,
	}
}

func SetErrorResponse(status, err, responseId string) globals.GeneralError {
	return globals.GeneralError{
		Status:     status,
		Error:      err,
		ResponseId: responseId,
	}
}
