package repositories

import (
	"errors"

	"github.com/hereswilson/jurassic-park-api/models"
	"gorm.io/gorm"
)

type CageRepository interface {
	CreateCage(*models.Cage) (*models.Cage, error)
	UpdateCage(*models.Cage) (*models.Cage, error)
	DeleteCageByName(string) error
	GetAllCages() ([]models.Cage, error)
	GetCageByName(string) (*models.Cage, error)
	GetCageByID(uint) (*models.Cage, error)
	GetCagesByPowerStatus(string) ([]models.Cage, error)
	GetDinosaurs(string) ([]models.Dinosaur, error)
	GetDinosaursBySpecies(string, string) ([]models.Dinosaur, error)
	GetNumDinosaurs(string) (int64, error)
	AddDinosaur(*models.Dinosaur) error
}

type cageRepository struct {
	db *gorm.DB
}

func NewCageRepository(db *gorm.DB) CageRepository {
	return &cageRepository{db: db}
}

func (c *cageRepository) GetCageByName(name string) (*models.Cage, error) {
	var cage models.Cage
	result := c.db.Where("name = ?", name)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, models.ErrCageNotFound
		}
		return nil, result.Error
	}
	return &cage, nil
}

func (c *cageRepository) CreateCage(cage *models.Cage) (*models.Cage, error) {
	result := c.db.Create(cage)
	if result.Error != nil {
		return nil, result.Error
	}
	return cage, nil
}

func (c *cageRepository) UpdateCage(cage *models.Cage) (*models.Cage, error) {
	result := c.db.Save(cage)
	if result.Error != nil {
		return nil, result.Error
	}
	return cage, nil
}

func (c *cageRepository) DeleteCageByName(name string) error {
	cage, err := c.GetCageByName(name)
	if err != nil {
		return err
	}
	result := c.db.Delete(cage)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *cageRepository) GetAllCages() ([]models.Cage, error) {
	var cages []models.Cage
	err := c.db.Find(&cages).Error
	return cages, err
}

func (c *cageRepository) GetCagesByPowerStatus(powerStatus string) ([]models.Cage, error) {
	var cages []models.Cage
	result := c.db.Where("power_status = ?", powerStatus).Find(&cages)
	if result.Error != nil {
		return nil, result.Error
	}
	return cages, nil
}

func (c *cageRepository) GetCageByID(id uint) (*models.Cage, error) {
	var cage models.Cage
	result := c.db.First(&cage, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, models.ErrCageNotFound
		}
		return nil, result.Error
	}
	return &cage, nil
}

func (c *cageRepository) GetDinosaurs(cageName string) ([]models.Dinosaur, error) {
	var cage models.Cage
	if err := c.db.Where("name = ?", cageName).Preload("Dinosaurs").First(&cage).Error; err != nil {
		return nil, err
	}
	return cage.Dinosaurs, nil
}

func (c *cageRepository) GetDinosaursBySpecies(cageName string, dinosaurSpecies string) ([]models.Dinosaur, error) {
	var cage models.Cage
	if err := c.db.Where("id = ?", cageName).Preload("Dinosaurs", "Species = ?", dinosaurSpecies).First(&cage).Error; err != nil {
		return nil, err
	}
	return cage.Dinosaurs, nil
}

func (c *cageRepository) GetNumDinosaurs(cageName string) (int64, error) {
	var count int64
	err := c.db.Model(&models.Dinosaur{}).Where("Name = ?", cageName).Count(&count).Error
	return count, err
}

func (c *cageRepository) AddDinosaur(dinosaur *models.Dinosaur) error {
	result := c.db.Create(dinosaur)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
