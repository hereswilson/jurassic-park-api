package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hereswilson/jurassic-park-api/models"
)

func GetDinosaurs(c *gin.Context) {
	dinosaurs, err := models.GetDinosaurs()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.IndentedJSON(http.StatusOK, dinosaurs)
}

func CreateDinosaur(c *gin.Context) {
	var dinosaur models.Dinosaur
	if err := c.ShouldBindJSON(&dinosaur); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	newDinosaur, err := models.CreateDinosaur(&dinosaur)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.IndentedJSON(http.StatusCreated, newDinosaur)
}

func GetDinosaurByName(c *gin.Context) {
	name := c.Param("name")
	dinosaur, err := models.GetDinosaurByName(name)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.IndentedJSON(http.StatusOK, dinosaur)
}

func UpdateDinosaur(c *gin.Context) {
	name := c.Param("name")
	var dinosaur models.Dinosaur
	if err := c.ShouldBindJSON(&dinosaur); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err := dinosaur.UpdateDinosaur(name)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.IndentedJSON(http.StatusOK, dinosaur)
}

func DeleteDinosaur(c *gin.Context) {
	name := c.Param("name")
	var dinosaur models.Dinosaur
	err := dinosaur.DeleteDinosaur(name)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Dinosaur deleted successfully!"})
}
