package core

import (
	"encoding/json"
	"net/http"
)

func WriteJSONResponse(
	w http.ResponseWriter,
	response interface{},
	status int,
) {
	w.Header().Set(
		"Content-Type",
		"application/json; charset=UTF-8",
	)
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
