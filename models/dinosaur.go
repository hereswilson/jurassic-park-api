package models

import "gorm.io/gorm"

type Dinosaur struct {
	gorm.Model
	Name      string `json:"name"`
	SpeciesID int
	Species   Species `json:"species"`
	CageID    int     `json:"cage_id"`
}
