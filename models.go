package main

import (
	_ "github.com/jinzhu/gorm"	
)

type Person struct {
	//gorm.Model
	ID             uint			   `gorm:"primary_key" json:"id"`
	Name           string		   `json:"name"`		
	City           string	       `json:"city"`
	Birthday       string	       `json:"birthday"`  
	ProfilePath    string		   `json:"profile_path"`
	SocialNetworks []SocialNetwork `gorm:"foreignkey:PersonID" json:"social_networks,omitempty"`
	Genres         []Genre         `gorm:"foreignkey:PersonID" json:"genres,omitempty"`
	Instruments    []Instrument    `gorm:"foreignkey:PersonID" json:"instruments,omitempty"`
}

type SocialNetwork struct {
	//gorm.Model
	ID       uint   `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	PersonID uint   `json:"person_id"`
}

type Genre struct {
	//gorm.Model
	ID       uint   `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	PersonID uint   `json:"person_id"`
}

type Instrument struct {
	//gorm.Model
	ID       uint   `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	PersonID uint   `json:"person_id"`
}
