package server

import (
	"UnnecessaryMafia-Backend/controller"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	controller.InsertUser(
		r.FormValue("username"),
		r.FormValue("password"),
		r.FormValue("phone_number"),
		r.FormValue("email"),
		r.FormValue("fname"),
		r.FormValue("lname"),
		r.FormValue("status"))
	http.Redirect(w, r, "/", 301)
}
