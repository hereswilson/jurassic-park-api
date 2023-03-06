package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hereswilson/jurassic-park-api/models"
	"github.com/hereswilson/jurassic-park-api/services"
)

type SpeciesController struct {
	speciesService *services.SpeciesService
}

func NewSpeciesController(speciesService *services.SpeciesService) *SpeciesController {
	return &SpeciesController{
		speciesService: speciesService,
	}
}

func (c *SpeciesController) GetSpecies(ctx *gin.Context) {
	species, err := c.speciesService.GetAllSpecies()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, species)
}

func (c *SpeciesController) CreateSpecies(ctx *gin.Context) {
	var species models.Species
	err := ctx.BindJSON(&species)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdSpecies, err := c.speciesService.CreateSpecies(&species)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdSpecies)
}

func (c *SpeciesController) GetSpecificSpecies(ctx *gin.Context) {
	speciesName := ctx.Param("species")

	species, err := c.speciesService.GetSpeciesByName(speciesName)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if species == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "species not found"})
		return
	}

	ctx.JSON(http.StatusOK, species)
}

func (c *SpeciesController) UpdateSpecies(ctx *gin.Context) {
	speciesName := ctx.Param("species")

	var species models.Species
	err := ctx.BindJSON(&species)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	species.Species = speciesName

	updatedSpecies, err := c.speciesService.UpdateSpecies(&species)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedSpecies)
}

func (c *SpeciesController) DeleteSpecies(ctx *gin.Context) {
	speciesName := ctx.Param("species")

	err := c.speciesService.DeleteSpecies(speciesName)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
