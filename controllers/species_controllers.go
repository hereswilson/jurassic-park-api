package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hereswilson/jurassic-park-api/models"
)

func GetSpecies(c *gin.Context) {
	species, err := models.GetSpecies()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.IndentedJSON(http.StatusOK, species)
}

func CreateSpecies(c *gin.Context) {
	var species models.Species
	if err := c.ShouldBindJSON(&species); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	newSpecies, err := models.CreateSpecies(&species)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.IndentedJSON(http.StatusCreated, newSpecies)
}

func GetSpecificSpecies(c *gin.Context) {
	name := c.Param("species")
	species, err := models.GetSpecificSpecies(name)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.IndentedJSON(http.StatusOK, species)
}

func UpdateSpecies(c *gin.Context) {
	specName := c.Param("species")
	var species models.Species
	if err := c.ShouldBindJSON(&species); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err := species.UpdateSpecies(specName)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.IndentedJSON(http.StatusOK, species)
}

func DeleteSpecies(c *gin.Context) {
	specName := c.Param("species")
	var species models.Species
	err := species.DeleteSpecies(specName)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Species deleted successfully!"})
}
