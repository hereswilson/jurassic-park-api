package models

import (
	"github.com/hereswilson/jurassic-park-api/database"
	"gorm.io/gorm"
)

type Species struct {
	gorm.Model
	Species string `json:"species" gorm:"type:text; not null; unique"`
	Diet    string `json:"diet" gorm:"type:text; not null"`
}

func GetSpecies() ([]Species, error) {
	var species []Species
	err := database.DB.Find(&species).Error
	return species, err
}

func GetSpecificSpecies(name string) (species Species, err error) {
	err = database.DB.Where("species = ?", name).First(&species).Error
	if err != nil {
		return Species{}, err
	}
	return species, nil
}

func CreateSpecies(species *Species) (*Species, error) {
	err := database.DB.Create(&species).Error
	if err != nil {
		return &Species{}, err
	}
	return species, nil
}

func (species *Species) UpdateSpecies(speciesName string) (err error) {

	err = database.DB.Where("species = ?", speciesName).Updates(&species).Error
	return err
}

func (species *Species) DeleteSpecies(speciesName string) (err error) {
	err = database.DB.Where("species = ?", speciesName).Delete(&species).Error
	return err
}
