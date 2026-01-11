package response

type ErrorSchema struct {
	Errors map[string]string `json:"errors"`
}

type ValidateError struct {
	Errors map[string]string `json:"errors"`
}

func ErrorResp(err error) *ErrorSchema {
	return &ErrorSchema{Errors: map[string]string{
		"message": err.Error(),
	}}
}
