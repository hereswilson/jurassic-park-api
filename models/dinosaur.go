package models

import (
	"gorm.io/gorm"
)

type Dinosaur struct {
	gorm.Model
	Name      string  `json:"name" gorm:"type:text; not null"`
	SpeciesID int     `json:"species_id" gorm:"type:int; not null"`
	Species   Species `json:"species" gorm:"not null"`
	CageID    uint    `json:"cage_id" gorm:"type:int; not null"`
}
