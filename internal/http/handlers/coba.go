package handlers

import (
	"net/http"

	"github.com/payfazz/go-handler"
	"github.com/payfazz/go-handler/defresponse"
)

// Profile .
type Profile struct {
	Name    string
	Hobbies []string
}

// Coba .
func Coba(r *http.Request) *handler.Response {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}

	// js, _ := json.Marshal(profile)
	return defresponse.JSON(200, profile) // we can't forget this, because it'll be compile error if there is no `return`
}
