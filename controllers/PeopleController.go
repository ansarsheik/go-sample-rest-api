package controllers

import (
        . "github.com/gorilla/mux"
	. "rest-api/entities"
	. "net/http"
	. "encoding/json"
)

type PeopleController interface {
	GetPeople(res ResponseWriter, req *Request)
	GetPerson(res ResponseWriter, req *Request)
	CreatePerson(res ResponseWriter, req *Request)
	DeletePerson(res ResponseWriter, req *Request)
}

type peopleController struct {
	People []Person
}

func NewPeopleController() PeopleController {
	peopleController := new(peopleController)

	var people []Person
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

	peopleController.People = people

	return peopleController
}

func (p peopleController) GetPeople(res ResponseWriter, req *Request) {
	NewEncoder(res).Encode(p.People)
}

func (p peopleController) GetPerson(res ResponseWriter, req *Request) {
	params := Vars(req)
    for _, item := range p.People {
        if item.ID == params["id"] {
            NewEncoder(res).Encode(item)
            return
        }
    }
    NewEncoder(res).Encode(&Person{})
}

func (p peopleController) CreatePerson(res ResponseWriter, req *Request) {
    params := Vars(req)
    var person Person
    _ = NewDecoder(req.Body).Decode(&person)
    person.ID = params["id"]
    p.People = append(p.People, person)
    NewEncoder(res).Encode(p.People)
}

func (p peopleController) DeletePerson(res ResponseWriter, req *Request) {
    params := Vars(req)
    for index, item := range p.People {
        if item.ID == params["id"] {
            p.People = append(p.People[:index], p.People[index+1:]...)
            break
        }
        NewEncoder(res).Encode(p.People)
    }
}
