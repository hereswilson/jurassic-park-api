package repositories

import (
	"errors"

	"github.com/hereswilson/jurassic-park-api/models"
	"gorm.io/gorm"
)

type DinosaurRepository interface {
	CreateDinosaur(dinosaur *models.Dinosaur) error
	UpdateDinosaur(dinosaur *models.Dinosaur) error
	DeleteDinosaurByName(name string) error
	GetDinosaurByName(name string) (*models.Dinosaur, error)
	GetDinosaursByCageID(cageID uint) ([]*models.Dinosaur, error)
	GetDinosaursBySpecies(species string) ([]*models.Dinosaur, error)
	GetAllDinosaurs() ([]*models.Dinosaur, error)
	GetDinosaursByCageName(name string) ([]*models.Dinosaur, error)
	IsSpeciesValid(species string) (bool, error)
}

type dinosaurRepository struct {
	db *gorm.DB
}

func NewDinosaurRepository(db *gorm.DB) DinosaurRepository {
	return &dinosaurRepository{db}
}

func (r *dinosaurRepository) CreateDinosaur(dinosaur *models.Dinosaur) error {
	return r.db.Create(dinosaur).Error
}

func (r *dinosaurRepository) UpdateDinosaur(dinosaur *models.Dinosaur) error {
	return r.db.Save(dinosaur).Error
}

func (r *dinosaurRepository) DeleteDinosaurByName(name string) error {
	return r.db.Where("name = ?", name).Delete(&models.Dinosaur{}).Error
}

func (r *dinosaurRepository) GetDinosaurByName(name string) (*models.Dinosaur, error) {
	var dinosaur models.Dinosaur
	err := r.db.Where("name = ?", name).First(&dinosaur).Error
	if err != nil {
		return nil, err
	}
	return &dinosaur, nil
}

func (r *dinosaurRepository) GetDinosaursByCageID(cageID uint) ([]*models.Dinosaur, error) {
	var dinosaurs []*models.Dinosaur
	err := r.db.Where("cage_id = ?", cageID).Find(&dinosaurs).Error
	if err != nil {
		return nil, err
	}
	return dinosaurs, nil
}

func (r *dinosaurRepository) GetAllDinosaurs() ([]*models.Dinosaur, error) {
	var dinosaurs []*models.Dinosaur
	err := r.db.Find(&dinosaurs).Error
	if err != nil {
		return nil, err
	}
	return dinosaurs, nil
}

func (r *dinosaurRepository) GetDinosaursByCageName(name string) ([]*models.Dinosaur, error) {
	var cage models.Cage
	err := r.db.Where("name = ?", name).First(&cage).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("cage not found")
		}
		return nil, err
	}

	var dinosaurs []*models.Dinosaur
	err = r.db.Where("cage_id = ?", cage.ID).Find(&dinosaurs).Error
	if err != nil {
		return nil, err
	}
	return dinosaurs, nil
}

func (r *dinosaurRepository) GetDinosaursBySpecies(species string) ([]*models.Dinosaur, error) {
	var dinosaurs []*models.Dinosaur
	err := r.db.Where("species = ?", species).Find(&dinosaurs).Error
	if err != nil {
		return nil, err
	}
	return dinosaurs, nil
}

func (r *dinosaurRepository) IsSpeciesValid(species string) (bool, error) {
	var count int64
	result := r.db.Model(&models.Species{}).Where("species = ?", species).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}
