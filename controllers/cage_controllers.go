package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hereswilson/jurassic-park-api/models"
)

func GetCages(c *gin.Context) {
	cages, err := models.GetCages()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.IndentedJSON(http.StatusOK, cages)
}

func CreateCage(c *gin.Context) {
	var cage models.Cage
	if err := c.ShouldBindJSON(&cage); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	newCage, err := models.CreateCage(&cage)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.IndentedJSON(http.StatusCreated, newCage)
}

func GetCageByName(c *gin.Context) {
	name := c.Param("name")
	cage, err := models.GetCageByName(name)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.IndentedJSON(http.StatusOK, cage)
}

func UpdateCage(c *gin.Context) {
	name := c.Param("name")
	var cage models.Cage
	if err := c.ShouldBindJSON(&cage); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err := cage.UpdateCage(name)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.IndentedJSON(http.StatusOK, cage)
}

func DeleteCage(c *gin.Context) {
	name := c.Param("name")
	var cage models.Cage
	err := cage.DeleteCage(name)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Cage deleted successfully!"})
}
