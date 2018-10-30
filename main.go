package main

import (
	"log"
	. "net/http"
    . "github.com/gorilla/mux"
    . "rest-api/controllers"
)

func main() {
    
    peopleController := NewPeopleController()

    router := NewRouter()
	router.HandleFunc("/people", peopleController.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", peopleController.GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", peopleController.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", peopleController.DeletePerson).Methods("DELETE")

	log.Fatal(ListenAndServe(":8000", router))
}