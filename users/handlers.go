package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ErrorOutput struct {
	message string
	errors  interface{}
}

func ValidateRequest(
	request *http.Request,
	validate *validator.Validate,
	data interface{},
) error {
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error while decoding the request")
		return err
	}

	err = validate.Struct(data)
	if err != nil {
		fmt.Println("Error while validating the request")
		return err
	}

	return nil
}

func WriteResponse(
	writer *http.ResponseWriter, 
	response interface{},
	status int,
) {
	// writer.WriteHeader(http.StatusBadRequest)
	(*writer).WriteHeader(status)
	json.NewEncoder(*writer).Encode(response)
}

func CreateListUsersHandler(db *gorm.DB, validate *validator.Validate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		defer r.Body.Close()

		var data UserCreateInput
		err := ValidateRequest(r, validate, &data)
		if err != nil {
			// TODO: make a switch to check the type of the error
			for _, error := range err.(validator.ValidationErrors) {
				fmt.Println(error.Field())
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user, err := CreateUser(db, &data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		response := UserCreateOutput{
			Email:     user.Email,
			FirstName: user.FirstName,
		}
		fmt.Println(response)
		WriteResponse(&w, &response, http.StatusOK)
	}
}

// func CreateListUserHandler(w http.ResponseWriter, r *http.Request) {
// 	// Handler logic here
// }
