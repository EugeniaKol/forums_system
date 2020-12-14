package forums

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/EugeniaKol/forums_system/server/tools"
)

// HTTPHandlerFunc create a var of its type
type HTTPHandlerFunc http.HandlerFunc

// HTTPHandler creates a new instance of channels HTTP handler.
func HTTPHandler(store *Store) HTTPHandlerFunc {
	log.Printf("Hello")
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("Listening")
		if r.Method == "GET" {
			log.Printf("Received get request")
			handleListForums(store, rw)
		} else if r.Method == "POST" {
			handleUserCreate(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleUserCreate(r *http.Request, rw http.ResponseWriter, store *Store) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		log.Printf("Error decoding channel input: %s", err)
		tools.WriteJSONBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.CreateUser(u.Nickname, u.Interests)
	if err == nil {
		tools.WriteJSONOk(rw, &u)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJSONInternalError(rw)
	}
}

func handleListForums(store *Store, rw http.ResponseWriter) {
	res, err := store.ListForums()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJSONInternalError(rw)
		return
	}
	tools.WriteJSONOk(rw, res)
}
