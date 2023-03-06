package services

import (
	"errors"

	"github.com/hereswilson/jurassic-park-api/models"
	"github.com/hereswilson/jurassic-park-api/repositories"
)

type SpeciesService struct {
	speciesRepo repositories.SpeciesRepository
}

func NewSpeciesService(speciesRepository repositories.SpeciesRepository) *SpeciesService {
	return &SpeciesService{
		speciesRepo: speciesRepository,
	}
}

func (s *SpeciesService) CreateSpecies(species *models.Species) (*models.Species, error) {
	if species == nil {
		return nil, errors.New("invalid species")
	}

	return s.speciesRepo.CreateSpecies(species)
}

func (s *SpeciesService) UpdateSpecies(species *models.Species) (*models.Species, error) {
	if species == nil {
		return nil, errors.New("invalid species")
	}

	existingSpecies, err := s.speciesRepo.FindSpeciesByName(species.Species)
	if err != nil {
		return nil, err
	}

	if existingSpecies == nil {
		return nil, errors.New("species not found")
	}

	return s.speciesRepo.UpdateSpecies(species)
}

func (s *SpeciesService) DeleteSpecies(speciesName string) error {
	if speciesName == "" {
		return errors.New("invalid species name")
	}

	existingSpecies, err := s.speciesRepo.FindSpeciesByName(speciesName)
	if err != nil {
		return err
	}

	if existingSpecies == nil {
		return errors.New("species not found")
	}

	return s.speciesRepo.DeleteSpecies(existingSpecies)
}

func (s *SpeciesService) GetSpeciesByName(speciesName string) (*models.Species, error) {
	if speciesName == "" {
		return nil, errors.New("invalid species name")
	}

	return s.speciesRepo.FindSpeciesByName(speciesName)
}

func (s *SpeciesService) GetAllSpecies() ([]*models.Species, error) {
	return s.speciesRepo.ListSpecies()
}
