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
				r.FormValue("status")) //ToDo Still accepts empty strings, only lname and status and phone can be empty
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
			_, err := w.Write([]byte("User Has Been Created!"))
			if err != nil {
				log.Printf("could not write response: %s\n", r.RequestURI)
			}
		},
		Catch: func(exception controller.Exception) {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(http.StatusInternalServerError)
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
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
	_, err = w.Write(jsonResp) // This way it doesn't superflow response call, ToDo should still test if it reaches this line
	if err != nil {
		log.Printf("could not write response: %s\n", r.RequestURI)
	}
}
