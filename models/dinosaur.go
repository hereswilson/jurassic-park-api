package models

import (
	"github.com/hereswilson/jurassic-park-api/database"
	"gorm.io/gorm"
)

type Dinosaur struct {
	gorm.Model
	Name      string  `json:"name" gorm:"type:text; not null"`
	SpeciesID int     `json:"species_id" gorm:"type:int; not null"`
	Species   Species `json:"species" gorm:"not null"`
	CageID    int     `json:"cage_id" gorm:"type:int; not null"`
}

func GetDinosaurs() (dinosaurs []Dinosaur, err error) {
	err = database.DB.Preload("Cage").Find(&dinosaurs).Error
	return dinosaurs, err
}

func GetDinosaursByCageID(cageID uint) (dinosaurs []Dinosaur, err error) {
	err = database.DB.Preload("Cage").Where("cage_id = ?", cageID).Find(&dinosaurs).Error
	return dinosaurs, err
}

func GetDinosaursBySpecies(species string) (dinosaurs []Dinosaur, err error) {
	err = database.DB.Preload("Cage").Where("species = ?", species).Find(&dinosaurs).Error
	return dinosaurs, err
}

func GetDinosaurByName(name string) (dinosaur Dinosaur, err error) {
	err = database.DB.Preload("Cage").Where("name = ?", name).First(&dinosaur).Error
	if err != nil {
		return Dinosaur{}, err
	}
	return dinosaur, nil
}

func CreateDinosaur(dinosaur *Dinosaur) (*Dinosaur, error) {
	err := database.DB.Create(&dinosaur).Error
	if err != nil {
		return &Dinosaur{}, err
	}
	return dinosaur, nil
}

func (dinosaur *Dinosaur) UpdateDinosaur(name string) (err error) {
	err = database.DB.Where("name = ?", name).Updates(&dinosaur).Error
	return err
}

func (dinosaur *Dinosaur) DeleteDinosaur(name string) (err error) {
	err = database.DB.Where("name = ?", name).Delete(&dinosaur).Error
	return err
}

func GetCageForDinosaur(name string) (*Cage, error) {
	var cage Cage
	if err := database.DB.Model(&Dinosaur{}).Joins("JOIN cage ON cage.id = dinosaurs.cage_id").Where("dinosaur.name = ?", name).Error; err != nil {
		return nil, err
	}
	return &cage, nil
}
