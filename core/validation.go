package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func RegisterValidatorJSONTag(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}

func NewValidate() *validator.Validate {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterTagNameFunc(RegisterValidatorJSONTag)
	return validate
}

func ValidateBody(
	r *http.Request,
	validate *validator.Validate,
	schema interface{},
	errorResponse *ErrorResponse,
) error {
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		fmt.Println("Error while decoding the request")
		return err
	}

	err = validate.Struct(schema)
	if err != nil {
		fmt.Println("Error while validating the request")
		for _, error := range err.(validator.ValidationErrors) {
			errorResponse.Add(
				FieldError{
					Field:   error.Field(),
					Message: error.Tag(),
				},
			)
		}
		return err
	}

	return nil
}
