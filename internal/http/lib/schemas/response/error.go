package response

type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type ValidateError struct {
	Errors map[string]string `json:"errors"`
}

func NewErrorResponse(message string) *Error {
	return &Error{
		Type:    "error",
		Message: message,
	}
}
