package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api/entities"

	"github.com/gorilla/mux"
)

// PeopleController manages operations about people
type PeopleController interface {
	GetPeople(res http.ResponseWriter, req *http.Request)
	GetPerson(res http.ResponseWriter, req *http.Request)
	CreatePerson(res http.ResponseWriter, req *http.Request)
	DeletePerson(res http.ResponseWriter, req *http.Request)
}

type peopleController struct {
	People []entities.Person
}

// NewPeopleController creates a new PeopleController
func NewPeopleController() PeopleController {
	peopleController := new(peopleController)

	var people []entities.Person
	people = append(people, entities.Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &entities.Address{City: "City X", State: "State X"}})
	people = append(people, entities.Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &entities.Address{City: "City Z", State: "State Y"}})
	people = append(people, entities.Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

	peopleController.People = people

	return peopleController
}

func (p peopleController) GetPeople(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(p.People)
}

func (p peopleController) GetPerson(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range p.People {
		if item.ID == params["id"] {
			json.NewEncoder(res).Encode(item)
			return
		}
	}
	json.NewEncoder(res).Encode(&entities.Person{})
}

func (p peopleController) CreatePerson(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person entities.Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	p.People = append(p.People, person)
	json.NewEncoder(res).Encode(p.People)
}

<<<<<<< HEAD
func (p peopleController) DeletePerson(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range p.People {
		if item.ID == params["id"] {
			p.People = append(p.People[:index], p.People[index+1:]...)
			break
		}
		json.NewEncoder(res).Encode(p.People)
	}
=======
func (p peopleController) DeletePerson(res ResponseWriter, req *Request) {
    params := Vars(req)
    for index, item := range p.People {
        if item.ID == params["id"] {
            p.People = append(p.People[:index], p.People[index+1:]...)
            break
        }
        NewEncoder(res).Encode(p.People)
    }
>>>>>>> refs/remotes/origin/master
}
