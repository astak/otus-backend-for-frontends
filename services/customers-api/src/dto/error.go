package dto

const (
	errorBadRequest int32 = 400
	errorNotFound   int32 = 404
)

type errorResponse struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func NewBadRequestError(message string) *errorResponse {
	return &errorResponse{
		Code:    errorBadRequest,
		Message: message,
	}
}

func NewNotFoundError(message string) *errorResponse {
	return &errorResponse{
		Code:    errorNotFound,
		Message: message,
	}
}

func NewNoAccountInfoError() *errorResponse {
	return &errorResponse {
		Code: errorBadRequest,
		Message: "Account info not available",
	}
}
