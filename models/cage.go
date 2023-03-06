package models

import (
	"errors"

	"gorm.io/gorm"
)

type Cage struct {
	gorm.Model
	Name            string     `json:"name" gorm:"not null;uniqueIndex"`
	MaximumCapacity int        `json:"maximum_capacity" gorm:"not null"`
	DinosaurCount   int        `json:"dinosaur_count" gorm:"not null;default:0"`
	PowerStatus     string     `json:"power_status" gorm:"not null;default:'ACTIVE'"`
	Dinosaurs       []Dinosaur `json:"dinosaurs" gorm:"foreignKey:CageID"`
}

var (
	ErrCageNotFound            = errors.New("cage not found")
	ErrCageNotEmpty            = errors.New("cage not empty")
	ErrDinosaurInDifferentCage = errors.New("dinosaur cannot be in this cages")
	ErrCagePoweredOn           = errors.New("cage powered on")
)
