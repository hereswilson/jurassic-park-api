package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hereswilson/jurassic-park-api/models"
	"github.com/hereswilson/jurassic-park-api/services"
)

type DinosaurController struct {
	dinoService *services.DinosaurService
}

func NewDinosaurController(dinoService *services.DinosaurService) *DinosaurController {
	return &DinosaurController{dinoService}
}

func (c *DinosaurController) GetDinosaurs(ctx *gin.Context) {
	dinosaurs, err := c.dinoService.GetAllDinosaurs()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, dinosaurs)
}

func (c *DinosaurController) CreateDinosaur(ctx *gin.Context) {
	var dinosaur models.Dinosaur
	err := ctx.ShouldBindJSON(&dinosaur)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cage, err := c.dinoService.CreateDinosaur(&dinosaur)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, cage)
}

func (c *DinosaurController) GetDinosaurByName(ctx *gin.Context) {
	name := ctx.Param("name")
	dinosaur, err := c.dinoService.GetDinosaurByName(name)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, dinosaur)
}

func (c *DinosaurController) UpdateDinosaur(ctx *gin.Context) {
	name := ctx.Param("name")
	var dinosaur models.Dinosaur
	err := ctx.ShouldBindJSON(&dinosaur)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dinosaur.Name = name
	err = c.dinoService.UpdateDinosaur(&dinosaur)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (c *DinosaurController) DeleteDinosaur(ctx *gin.Context) {
	name := ctx.Param("name")
	err := c.dinoService.DeleteDinosaurByName(name)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (c *DinosaurController) AddDinosaurToCage(ctx *gin.Context) {
	name := ctx.Param("name")
	cageName := ctx.Param("cage_name")
	dinosaur, err := c.dinoService.AddDinosaurToCage(name, cageName)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, dinosaur)
}

func (c *DinosaurController) RemoveDinosaurFromCage(ctx *gin.Context) {
	name := ctx.Param("name")
	err := c.dinoService.RemoveDinosaurFromCage(name)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (c *DinosaurController) FilterDinosaursBySpecies(ctx *gin.Context) {
	species := ctx.Param("species")
	dinosaurs, err := c.dinoService.GetDinosaursBySpecies(species)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, dinosaurs)
}
