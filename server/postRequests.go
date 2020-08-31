package server

import (
	"UnnecessaryMafia-Backend/controller"
	"encoding/json"
	"log"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	controller.Block{
		Try: func() {
			controller.InsertUser(
				r.FormValue("username"),
				r.FormValue("password"),
				r.FormValue("phone_number"),
				r.FormValue("email"),
				r.FormValue("fname"),
				r.FormValue("lname"),
				r.FormValue("status"))
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		},
		Catch: func(exception controller.Exception) {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal(http.StatusInternalServerError)
		},
	}.Do()
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	gameUser := controller.GetUser(
		r.FormValue("username"),
		r.FormValue("password"))
	jsonResp, err := json.MarshalIndent(gameUser, "", "	")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if gameUser.Username == "" {
		jsonResp = []byte("Operation Failed")
	}
	_, err = w.Write(jsonResp) //ToDo  http: superfluous response.WriteHeader call is because of calling both 'write' and 'redirect'
	if err != nil {
		log.Printf("could not write response: %s", r.RequestURI)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
