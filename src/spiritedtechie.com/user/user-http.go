package main

import (
	"encoding/json"
	"fmt"
	"log"
	"spiritedtechie.com/user/models"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/user/{id:[0-9]+}", GetUserHandler).
	  Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	u := models.User{
		Name:   "John Smith",
		Gender: "male",
		Age :    50,
		Id:     id,
	}

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}
