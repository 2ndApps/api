package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)


func ListPerson(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=utf8")

	var person []Person
	db.Find(&person)

	err = json.NewEncoder(w).Encode(person)
	if err != nil {
		w.Write([]byte("could not parse persons to json"))
	}

}


func GetPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf8")

	params := mux.Vars(r)
	id := params["id"]
	
	var person Person
	if db.First(&person, id).RecordNotFound(){
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	db.Find(&person).Related(&person.SocialNetworks)
	db.Find(&person).Related(&person.Genres)
	db.Find(&person).Related(&person.Instruments)
	
	err = json.NewEncoder(w).Encode(person)
	if err != nil {
		w.Write([]byte("Could not parse person to json"))
	}
}


func CreatePerson(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	var person Person
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	db.Create(&person)
	w.WriteHeader(http.StatusCreated)
}


func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	db.Where("ID = ?", id).Delete(Person{})
}