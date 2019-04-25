package controllers

import (
	"net/http"

	u "../utils"
)

var Health = func(w http.ResponseWriter, r *http.Request) {
	response := u.Message(true, "SYSTEM IS not DOWN")
	u.Respond(w, response)
}
