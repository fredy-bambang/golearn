package route

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HandleRequests .
func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", Get).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

// Get .
func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from ExampleHandlerv1.")
}
