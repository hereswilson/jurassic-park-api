package models

import (
	"errors"
	"strings"

	"github.com/hereswilson/jurassic-park-api/database"
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

func CreateCage(cage *Cage) error {
	if err := database.DB.Create(&cage).Error; err != nil {
		return err
	}
	return nil
}

func GetCages() ([]Cage, error) {
	var cages []Cage
	if err := database.DB.Find(&cages).Error; err != nil {
		return nil, err
	}
	return cages, nil
}

func GetCageByName(name string) (*Cage, error) {
	var cage Cage
	if err := database.DB.Where("name = ?", name).Preload("Dinosaurs").First(&cage).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrCageNotFound
		}
		return nil, err
	}
	return &cage, nil
}

func (cage *Cage) UpdateCage(name string) error {
	result := database.DB.Where("name = ?", name).Updates(&cage)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (cage *Cage) DeleteCage(name string) error {
	var count int64
	if err := database.DB.Model(&Dinosaur{}).Joins("JOIN cage ON cage.id = dinosaurs.cage_id").Where("cage.name = ?", name).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return ErrCageNotEmpty
	}
	if cage.PowerStatus == "ACTIVE" {
		return ErrCagePoweredOn
	}
	if err := database.DB.Delete(&cage).Error; err != nil {
		return err
	}
	return nil
}

func GetDinosaursInCage(cageName string) (dinosaurs []Dinosaur, err error) {
	err = database.DB.Model(&Cage{}).
		Joins("LEFT JOIN dinosaurs ON cages.id = dinosaurs.cage_id").
		Where("cages.name = ?", cageName).
		Find(&dinosaurs).Error
	return dinosaurs, err
}

func FilterCagesByPowerStatus(powerStatus string) (cages []Cage, err error) {
	if powerStatus != "" {
		err = database.DB.Where("power_status = ?", powerStatus).Find(&cages).Error
	} else {
		err = database.DB.Find(&cages).Error
	}
	return cages, err
}

func (cage *Cage) AddDinosaur(dinosaur *Dinosaur) (err error) {
	if cage.DinosaurCount >= cage.MaximumCapacity {
		return errors.New("cage is already full")
	}
	if strings.ToLower(dinosaur.Species.Diet) == "carnivore" {
		for _, d := range cage.Dinosaurs {
			if d.Species != dinosaur.Species {
				return ErrDinosaurInDifferentCage
			}
		}
	} else {
		for _, d := range cage.Dinosaurs {
			if strings.ToLower(d.Species.Diet) == "carnivore" {
				return ErrDinosaurInDifferentCage
			}
		}
	}
	db := database.DB
	if err := db.Model(cage).Association("Dinosaurs").Append(dinosaur); err != nil {
		return err
	}

	cage.DinosaurCount++
	if err := db.Save(cage).Error; err != nil {
		return err
	}

	return nil
}

func (cage *Cage) RemoveDinosaur(dinosaur *Dinosaur) (err error) {
	if dinosaur.CageID != int(cage.ID) {
		return errors.New("dinosaur is not in this cage")
	}
	err = database.DB.Delete(&dinosaur).Error
	if err != nil {
		return err
	}
	cage.DinosaurCount--
	err = database.DB.Save(&cage).Error
	if err != nil {
		return err
	}
	return nil
}
