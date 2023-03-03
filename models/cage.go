package models

import "gorm.io/gorm"

type Cage struct {
	gorm.Model
	Name            string     `json:"name"`
	MaximumCapacity int        `json:"maximum_capacity"`
	DinosaurCount   int        `json:"dinosaur_count"`
	PowerStatus     bool       `json:"power_status"`
	Dinosaurs       []Dinosaur `json:"dinosaurs"`
}
