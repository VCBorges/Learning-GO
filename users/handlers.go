package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ErrorOutput struct {
	Message string
	Errors  interface{}
}

func ValidateRequest(
	r *http.Request,
	validate *validator.Validate,
	data interface{},
) error {
	err := json.NewDecoder(r.Body).Decode(&data)
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

func WriteJSONResponse(
	w http.ResponseWriter,
	response interface{},
	status int,
) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	if response == nil {
		response = make(map[string]string)
	}
	// TODO: handle more cases where the type of response is not a pointer to a struct

	json.NewEncoder(w).Encode(response)
}

func SuccessJSONResponse(
	w http.ResponseWriter,
	response interface{},
) {
	WriteJSONResponse(w, &response, http.StatusOK)
}

func ErrorJSONResponse(
	w http.ResponseWriter,
	err interface{},
) {
	WriteJSONResponse(w, err, http.StatusBadRequest)
}

func CreateListUsersHandler(db *gorm.DB, validate *validator.Validate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var data UserCreateInput
		err := ValidateRequest(r, validate, &data)
		if err != nil {
			// TODO: make a switch to check the type of the error
			for _, error := range err.(validator.ValidationErrors) {
				fmt.Println(error.Field())
			}
			ErrorJSONResponse(w, nil)
			return
		}

		_, err = GetUserByEmail(data.Email, db)
		if err == nil {
			fmt.Println("User already exists")
			ErrorJSONResponse(w, nil)
			return
		}

		user, err := CreateUser(db, &data)
		if err != nil {
			fmt.Println("Error while creating user")
			ErrorJSONResponse(w, nil)
			return
		}

		response := UserCreateOutput{
			Email:     user.Email,
			FirstName: user.FirstName,
		}
		fmt.Println(response)
		SuccessJSONResponse(w, &response)
	}
}
