package repositories

import (
	"errors"

	"github.com/hereswilson/jurassic-park-api/models"
	"gorm.io/gorm"
)

type SpeciesRepository interface {
	CreateSpecies(species *models.Species) (*models.Species, error)
	UpdateSpecies(species *models.Species) (*models.Species, error)
	DeleteSpecies(species *models.Species) error
	FindSpeciesByName(speciesName string) (*models.Species, error)
	ListSpecies() ([]*models.Species, error)
}

type speciesRepo struct {
	db *gorm.DB
}

func NewSpeciesRepository(db *gorm.DB) SpeciesRepository {
	return &speciesRepo{db: db}
}

func (r *speciesRepo) CreateSpecies(species *models.Species) (*models.Species, error) {
	if species == nil {
		return nil, errors.New("invalid species")
	}

	if result := r.db.Create(&species); result.Error != nil {
		return nil, result.Error
	}

	return species, nil
}

func (r *speciesRepo) UpdateSpecies(species *models.Species) (*models.Species, error) {
	if species == nil {
		return nil, errors.New("invalid species")
	}

	if result := r.db.Save(&species); result.Error != nil {
		return nil, result.Error
	}

	return species, nil
}

func (r *speciesRepo) DeleteSpecies(species *models.Species) error {
	if species == nil {
		return errors.New("invalid species")
	}

	if result := r.db.Delete(&species); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *speciesRepo) FindSpeciesByName(speciesName string) (*models.Species, error) {
	var species models.Species
	if result := r.db.Where("species = ?", speciesName).First(&species); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &species, nil
}

func (r *speciesRepo) ListSpecies() ([]*models.Species, error) {
	var speciesList []*models.Species
	if result := r.db.Find(&speciesList); result.Error != nil {
		return nil, result.Error
	}

	return speciesList, nil
}
