package models

import "gorm.io/gorm"

type Cage struct {
	gorm.Model
	Name            string     `json:"name" gorm:"type:text not null"`
	MaximumCapacity int        `json:"maximum_capacity" gorm:"type:int not null"`
	DinosaurCount   int        `json:"dinosaur_count" gorm:"type:int not null"`
	PowerStatus     string     `json:"power_status" gorm:"type:text not null"`
	Dinosaurs       []Dinosaur `json:"dinosaurs"`
}
