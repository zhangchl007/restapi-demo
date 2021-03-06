package main

import (
	"encoding/json"
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

type Person struct {
	ID string `json:"id,omitempty"`
	Firstname string `json:"firstname,zhang"`
	Lastname string `json:"lastname,jimmy"`
	Address *Address `json:"address,omitempty"`
}

type Address struct {
	City string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var  people []Person

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params :=mux.Vars(req)
	for _, item :=range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}


func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
    _ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people,person)
	json.NewEncoder(w).Encode(people)
}
func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params :=mux.Vars(req)
	for index,item :=range people {
		if item.ID == params["id"] {
			people =append(people[:index],people[index+1:]...)
		break
		}
	}

}

func main() {
	people = append(people,Person{ID: "1",Firstname: "zhang",Lastname: "jimmy" , Address: &Address{City: "shenzhen",State: "Guangdong"}})
	people = append(people,Person{ID: "2",Firstname: "liu",Lastname: "tony" , Address: &Address{City: "shenzhen",State: "Guangdong"}})
	r := mux.NewRouter()
	fmt.Println("HTTP server is started")
	r.HandleFunc("/people",GetPeopleEndpoint).Methods("GET")
	r.HandleFunc("/people/{id}",GetPersonEndpoint).Methods("GET")
	r.HandleFunc("/people/{id}",CreatePersonEndpoint).Methods("POST")
	r.HandleFunc("/people/{id}",DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080",r))

}
