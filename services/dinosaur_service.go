package services

import (
	"errors"

	"github.com/hereswilson/jurassic-park-api/models"
	"github.com/hereswilson/jurassic-park-api/repositories"
)

type DinosaurService struct {
	dinoRepo    repositories.DinosaurRepository
	cageRepo    repositories.CageRepository
	speciesRepo repositories.SpeciesRepository
}

func NewDinosaurService(dinosaurRepository repositories.DinosaurRepository, cageRepository repositories.CageRepository, speciesRepository repositories.SpeciesRepository) *DinosaurService {
	return &DinosaurService{
		dinoRepo:    dinosaurRepository,
		cageRepo:    cageRepository,
		speciesRepo: speciesRepository,
	}
}

func (s *DinosaurService) GetDinosaurByName(name string) (*models.Dinosaur, error) {
	return s.dinoRepo.GetDinosaurByName(name)
}

func (s *DinosaurService) CreateDinosaur(dinosaur *models.Dinosaur) (*models.Cage, error) {
	if dinosaur.Name == "" {
		return nil, errors.New("dinosaur name is required")
	}

	valid, err := s.dinoRepo.IsSpeciesValid(dinosaur.Species.Species)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, errors.New("invalid species")
	}

	cage, err := s.cageRepo.GetCageByID(dinosaur.CageID)
	if err != nil {
		return nil, err
	}

	if cage.PowerStatus != "ACTIVE" {
		return nil, errors.New("cannot add dinosaur to inactive cage")
	}

	if dinosaur.Species.Diet == "CARNIVORE" {
		dinosaursInCage, err := s.dinoRepo.GetDinosaursByCageID(dinosaur.CageID)
		if err != nil {
			return nil, err
		}

		for _, d := range dinosaursInCage {
			if d.Species != dinosaur.Species {
				return nil, errors.New("carnivores can only be in a cage with other dinosaurs of the same species")
			}
		}
	}

	if len(cage.Dinosaurs) >= cage.MaximumCapacity {
		return nil, errors.New("cage is at maximum capacity")
	}

	cage.DinosaurCount++
	cage.Dinosaurs = append(cage.Dinosaurs, *dinosaur)

	err = s.dinoRepo.CreateDinosaur(dinosaur)
	if err != nil {
		return nil, err
	}

	updatedCage, err := s.cageRepo.UpdateCage(cage)
	if err != nil {
		return nil, err
	}

	return updatedCage, nil
}

func (s *DinosaurService) UpdateDinosaur(dinosaur *models.Dinosaur) error {
	// Check if the species is valid
	valid, err := s.dinoRepo.IsSpeciesValid(dinosaur.Species.Species)

	if err != nil {
		return err
	}

	if !valid {
		return errors.New("invalid species")
	}

	// Check if the cage exists
	if _, err := s.cageRepo.GetCageByID(dinosaur.CageID); err != nil {
		return errors.New("cage not found")
	}

	return s.dinoRepo.UpdateDinosaur(dinosaur)
}

func (s *DinosaurService) DeleteDinosaurByName(name string) error {
	return s.dinoRepo.DeleteDinosaurByName(name)
}

func (s *DinosaurService) GetDinosaursByCageID(cageID uint) ([]*models.Dinosaur, error) {
	return s.dinoRepo.GetDinosaursByCageID(cageID)
}

func (s *DinosaurService) GetAllDinosaurs() ([]*models.Dinosaur, error) {
	return s.dinoRepo.GetAllDinosaurs()
}

func (s *DinosaurService) GetDinosaursByCageName(name string) ([]*models.Dinosaur, error) {
	cage, err := s.cageRepo.GetCageByName(name)
	if err != nil {
		return nil, err
	}

	return s.dinoRepo.GetDinosaursByCageID(cage.ID)
}

func (s *DinosaurService) GetDinosaursBySpecies(species string) ([]*models.Dinosaur, error) {
	return s.dinoRepo.GetDinosaursBySpecies(species)
}

func (s *DinosaurService) AddDinosaurToCage(name, cageName string) (*models.Dinosaur, error) {
	// First, find the dinosaur by name
	dinosaur, err := s.dinoRepo.GetDinosaurByName(name)
	if err != nil {
		return nil, err
	}

	if dinosaur == nil {
		return nil, errors.New("dinosaur not found")
	}

	// Next, find the cage by name
	cage, err := s.cageRepo.GetCageByName(cageName)
	if err != nil {
		return nil, err
	}

	if cage == nil {
		return nil, errors.New("cage not found")
	}

	if cage.PowerStatus != "ACTIVE" {
		return nil, errors.New("cannot add dinosaur to inactive cage")
	}

	if dinosaur.Species.Diet == "CARNIVORE" {
		dinosaursInCage, err := s.dinoRepo.GetDinosaursByCageID(dinosaur.CageID)
		if err != nil {
			return nil, err
		}

		for _, d := range dinosaursInCage {
			if d.Species != dinosaur.Species {
				return nil, errors.New("carnivores can only be in a cage with other dinosaurs of the same species")
			}
		}
	}

	if len(cage.Dinosaurs) >= cage.MaximumCapacity {
		return nil, errors.New("cage is at maximum capacity")
	}

	cage.DinosaurCount++
	dinosaur.CageID = cage.ID

	err = s.UpdateDinosaur(dinosaur)
	if err != nil {
		return nil, err
	}

	err = s.cageRepo.AddDinosaur(cage, dinosaur)
	if err != nil {
		return nil, err
	}

	_, err = s.cageRepo.UpdateCage(cage)
	if err != nil {
		return nil, err
	}

	return dinosaur, nil
}

func (s *DinosaurService) RemoveDinosaurFromCage(name string) error {
	dinosaur, err := s.dinoRepo.GetDinosaurByName(name)
	if err != nil {
		return err
	}

	if dinosaur == nil {
		return errors.New("dinosaur not found")
	}

	cage, err := s.cageRepo.GetCageByID(dinosaur.CageID)
	if err != nil {
		return err
	}

	if cage == nil {
		return errors.New("cage not found")
	}

	cage.DinosaurCount--
	dinosaur.CageID = 0

	err = s.UpdateDinosaur(dinosaur)
	if err != nil {
		return err
	}

	err = s.cageRepo.RemoveDinosaur(cage, dinosaur)
	if err != nil {
		return err
	}

	_, err = s.cageRepo.UpdateCage(cage)
	if err != nil {
		return err
	}

	return nil
}
