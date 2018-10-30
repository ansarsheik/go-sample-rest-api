package main

import (
	"log"
	"net/http"
	"rest-api/controllers"

	"github.com/gorilla/mux"
)

func main() {

	peopleController := controllers.NewPeopleController()

	router := mux.NewRouter()
	router.HandleFunc("/people", peopleController.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", peopleController.GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", peopleController.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", peopleController.DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
