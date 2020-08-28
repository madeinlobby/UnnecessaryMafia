package server

import (
	"UnnecessaryMafia-Backened/controller"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	password := r.FormValue("password")
	controller.InsertUser(name, password)
	http.Redirect(w, r, "/", 301)
}
