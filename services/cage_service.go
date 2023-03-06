package services

import (
	"errors"

	"github.com/hereswilson/jurassic-park-api/models"
	"github.com/hereswilson/jurassic-park-api/repositories"
)

type CageService struct {
	cageRepo repositories.CageRepository
}

func NewCageService(cageRepository repositories.CageRepository) *CageService {
	return &CageService{
		cageRepo: cageRepository,
	}
}

func (c *CageService) GetCageByName(name string) (*models.Cage, error) {
	cage, err := c.cageRepo.GetCageByName(name)
	if err != nil {
		return nil, err
	}

	if cage == nil {
		return nil, errors.New("cage not found")
	}

	return cage, nil
}

func (c *CageService) GetCages() ([]models.Cage, error) {
	return c.cageRepo.GetAllCages()
}

func (c *CageService) CreateCage(cage *models.Cage) (*models.Cage, error) {
	return c.cageRepo.CreateCage(cage)
}

func (c *CageService) UpdateCageByName(cage *models.Cage) (*models.Cage, error) {
	existingCage, err := c.cageRepo.GetCageByName(cage.Name)
	if err != nil {
		return nil, err
	}

	if existingCage == nil {
		return nil, errors.New("cage not found")
	}

	if cage.DinosaurCount > cage.MaximumCapacity || cage.MaximumCapacity < existingCage.DinosaurCount {
		return nil, errors.New("cannot update a cage with more dinosaurs than its capacity")
	}

	if cage.PowerStatus == "DOWN" && cage.DinosaurCount > 0 {
		return nil, errors.New("cannot power down a cage with dinosaurs inside")
	}

	return c.cageRepo.UpdateCage(cage)
}

func (c *CageService) DeleteCageByName(name string) error {
	existingCage, err := c.cageRepo.GetCageByName(name)
	if err != nil {
		return err
	}

	if existingCage == nil {
		return errors.New("cage not found")
	}

	if existingCage.DinosaurCount > 0 {
		return errors.New("cannot delete a cage with dinosaurs in it")
	}

	return c.cageRepo.DeleteCageByName(name)
}

func (c *CageService) GetDinosaursInCage(name string) ([]models.Dinosaur, error) {
	cage, err := c.GetCageByName(name)
	if err != nil {
		return nil, err
	}

	if cage == nil {
		return nil, errors.New("cage not found")
	}

	return c.cageRepo.GetDinosaurs(name)
}
