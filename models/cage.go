package models

import (
	"github.com/hereswilson/jurassic-park-api/database"
	"gorm.io/gorm"
)

type Cage struct {
	gorm.Model
	Name            string     `json:"name" gorm:"type:text not null"`
	MaximumCapacity int        `json:"maximum_capacity" gorm:"type:int not null"`
	DinosaurCount   int        `json:"dinosaur_count" gorm:"type:int not null"`
	PowerStatus     string     `json:"power_status" gorm:"type:text not null"`
	Dinosaurs       []Dinosaur `json:"dinosaurs"`
}

func CreateCage(cage *Cage) (*Cage, error) {
	err := database.DB.Create(&cage).Error
	if err != nil {
		return &Cage{}, err
	}
	return cage, nil
}

func GetCages() (cages []Cage, err error) {
	err = database.DB.Find(&cages).Error
	return cages, err
}

func GetCageByName(name string) (cage Cage, err error) {
	err = database.DB.Where("name = ?", name).First(&cage).Error
	if err != nil {
		return Cage{}, err
	}
	return cage, nil
}

func (cage *Cage) UpdateCage(name string) (err error) {
	err = database.DB.Where("name = ?", name).Updates(&cage).Error
	return err
}

func (cage *Cage) DeleteCage(name string) (err error) {
	err = database.DB.Where("name = ?", name).Delete(&cage).Error
	return err
}
