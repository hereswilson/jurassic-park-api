package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hereswilson/jurassic-park-api/models"
)

// Error responses
var (
	ErrNotFound   = gin.H{"error": "Resource not found"}
	ErrBadRequest = gin.H{"error": "Bad request"}
	ErrInternal   = gin.H{"error": "Internal server error"}
)

type CageController struct{}

func (ctrl *CageController) GetCages(c *gin.Context) {
	cages, err := models.GetCages()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.IndentedJSON(http.StatusOK, cages)
}

func (ctrl *CageController) CreateCage(c *gin.Context) {
	var cage models.Cage
	if err := c.ShouldBindJSON(&cage); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := models.CreateCage(&cage)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusCreated, cage)
}

func (ctrl *CageController) GetCageByName(c *gin.Context) {
	name := c.Param("name")
	cage, err := models.GetCageByName(name)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.IndentedJSON(http.StatusOK, cage)
}

func (ctrl *CageController) UpdateCage(c *gin.Context) {
	name := c.Param("name")
	var cage models.Cage
	if err := c.ShouldBindJSON(&cage); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := cage.UpdateCage(name)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusOK, cage)
}

func (ctrl *CageController) DeleteCage(c *gin.Context) {
	name := c.Param("name")
	var cage models.Cage
	err := cage.DeleteCage(name)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Cage deleted successfully!"})
}

func (ctrl *CageController) GetDinosaursInCage(c *gin.Context) {
	name := c.Param("name")
	dinosaurs, err := models.GetDinosaursInCage(name)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.IndentedJSON(http.StatusOK, dinosaurs)
}

func (ctrl *CageController) FilterCagesByPowerStatus(c *gin.Context) {
	powerStatus := c.Query("power_status")
	cages, err := models.FilterCagesByPowerStatus(powerStatus)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusOK, cages)
}
