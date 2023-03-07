package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hereswilson/jurassic-park-api/models"
	"github.com/hereswilson/jurassic-park-api/services"
)

type CageController struct {
	cageService *services.CageService
}

func NewCageController(cageService *services.CageService) *CageController {
	return &CageController{cageService: cageService}
}

func (c *CageController) GetCages(ctx *gin.Context) {
	cages, err := c.cageService.GetCages()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, cages)
}

func (c *CageController) CreateCage(ctx *gin.Context) {
	var cage models.Cage
	if err := ctx.ShouldBindJSON(&cage); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newCage, err := c.cageService.CreateCage(&cage)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, newCage)
}

func (c *CageController) GetCageByName(ctx *gin.Context) {
	cageName := ctx.Query("name")
	cage, err := c.cageService.GetCageByName(cageName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if cage == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "cage not found"})
		return
	}
	ctx.JSON(http.StatusOK, cage)
}

func (c *CageController) UpdateCage(ctx *gin.Context) {
	var cage models.Cage
	if err := ctx.ShouldBindJSON(&cage); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCage, err := c.cageService.UpdateCageByName(&cage)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedCage)
}

func (c *CageController) DeleteCage(ctx *gin.Context) {
	cageName := ctx.Query("name")
	err := c.cageService.DeleteCageByName(cageName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *CageController) GetDinosaursInCage(ctx *gin.Context) {
	cageName := ctx.Query("name")

	dinosaurs, err := c.cageService.GetDinosaursInCage(cageName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dinosaurs)
}
