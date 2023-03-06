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

func (cage *Cage) BeforeDelete(tx *gorm.DB) (err error) {
	// Check if cage is powered on before allowing deletion
	if cage.PowerStatus != "DOWN" {
		return errors.New("cannot delete a cage that is not powered down")
	}

	return nil
}

func (cage *Cage) BeforeUpdate(tx *gorm.DB) (err error) {
	// Check if cage is powered down before allowing power status update
	if tx.Statement.Changed("PowerStatus") && cage.PowerStatus != "DOWN" {
		return errors.New("cannot power off a cage that contains dinosaurs")
	}

	return nil
}
