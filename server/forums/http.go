package forums

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/EugeniaKol/forums_system/server/tools"
)

// HTTPHandlerFunc creates a var of its type
type HTTPHandlerFunc http.HandlerFunc

// HTTPHandler creates a new instance of channels HTTP handler.
func HTTPHandler(store *Store) HTTPHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("Listening...")
		if r.Method == "GET" {
			log.Printf("Received GET request")
			handleListForums(store, rw)
		} else if r.Method == "POST" {
			log.Printf("Received POST request")
			handleUserCreate(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleUserCreate(r *http.Request, rw http.ResponseWriter, store *Store) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		fmt.Printf("%+v\n", u)
		log.Printf("Error decoding channel input: %s", err)
		tools.WriteJSONBadRequest(rw, "Bad JSON payload")
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
