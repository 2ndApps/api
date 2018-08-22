package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func main() {
	router := mux.NewRouter()
	
	db, err = gorm.Open("sqlite3", "development.db")
	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()
	db.AutoMigrate(&Person{})
	db.AutoMigrate(&SocialNetwork{})
	db.AutoMigrate(&Genre{})
	db.AutoMigrate(&Instrument{})

	// Seeds
	person := &Person{
		Name:        "Ricardo",
		City:        "Matão",
		Birthday:    "14/05/1993",
		ProfilePath: "/EfnlUjNvfdpogHtfVcdDsa.jpg",
		SocialNetworks: []SocialNetwork{
			{Name: "Facebook", Username: "mattosri"},
			{Name: "Whatsapp", Username: "16993872222"},
		},
		Genres: []Genre{
			{Name: "Rock"},
			{Name: "Alternativo"},
			{Name: "Pop-Rock"},
		},
		Instruments: []Instrument{
			{Name: "Violao"},
		},
	}

	db.Create(person)

	router.HandleFunc("/persons", ListPerson).Methods("GET")
	router.HandleFunc("/persons", CreatePerson).Methods("POST")
	router.HandleFunc("/persons/{id:[0-9]+}", GetPerson).Methods("GET")
	router.HandleFunc("/persons/{id:[0-9+]}", DeletePerson).Methods("DELETE")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
