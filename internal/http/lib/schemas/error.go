package schemas

type ErrorResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type ValidateErrorResponse struct {
	Errors map[string]string `json:"errors"`
}

func NewErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{
		Type:    "error",
		Message: message,
	}
}
