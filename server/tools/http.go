package tools

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorObject struct {
	Message string `json:"message"`
}

//WriteJSONBadRequest sends 400 error response with a JSON object describing the error reason.
func WriteJSONBadRequest(rw http.ResponseWriter, message string) {
	writeJSON(rw, http.StatusBadRequest, &errorObject{Message: message})
}

//WriteJSONInternalError sends 500 error response.
func WriteJSONInternalError(rw http.ResponseWriter) {
	writeJSON(rw, http.StatusBadRequest, &errorObject{Message: "internal error happened"})
}

//WriteJSONOk sends 200 response to the client serializing the input object in JSON format.
func WriteJSONOk(rw http.ResponseWriter, res interface{}) {
	writeJSON(rw, http.StatusOK, res)
}

func writeJSON(rw http.ResponseWriter, status int, res interface{}) {
	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(status)
	err := json.NewEncoder(rw).Encode(res)
	if err != nil {
		log.Printf("Error writing response: %s", err)
	}
}
