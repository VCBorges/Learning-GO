package core

type ErrorResponse struct {
	Errors []FieldError `json:"errors"`
}

func (e *ErrorResponse) Add(err FieldError) {
	e.Errors = append(e.Errors, err)
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}