package main

import (
	"net/http"
	"project_name/database"
	"project_name/users"
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

func main() {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterTagNameFunc(RegisterValidatorJSONTag)

	db, err := database.InitDB()
	if err != nil {
		panic("failed to connect database")
	}
	// http.NewServeMux()
	http.HandleFunc("/users", users.CreateListUsersHandler(db, validate))
	// http.HandleFunc("/goodbye", goodbyeHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error to serve the HTTP server")
	}
}
