package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hereswilson/jurassic-park-api/models"
)

type DinosaurController struct{}

// GetDinosaursInCage returns a list of dinosaurs in a specific cage
func (dc *DinosaurController) GetDinosaursInCage(c *gin.Context) {
	cageName := c.Param("cageName")

	dinosaurs, err := models.GetDinosaursInCage(cageName)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, dinosaurs)
}

// GetDinosaurs returns a list of all dinosaurs
func (dc *DinosaurController) GetDinosaurs(c *gin.Context) {
	dinosaurs, err := models.GetDinosaurs()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.IndentedJSON(http.StatusOK, dinosaurs)
}

// GetDinosaur returns a single dinosaur by name
func (dc *DinosaurController) GetDinosaurByName(c *gin.Context) {
	name := c.Param("name")

	dinosaur, err := models.GetDinosaurByName(name)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, dinosaur)
}

// CreateDinosaur creates a new dinosaur
func (dc *DinosaurController) CreateDinosaur(c *gin.Context) {
	var dinosaur models.Dinosaur
	if err := c.ShouldBindJSON(&dinosaur); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	newDinosaur, err := models.CreateDinosaur(&dinosaur)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.IndentedJSON(http.StatusCreated, newDinosaur)
}

// UpdateDinosaur updates an existing dinosaur by name
func (dc *DinosaurController) UpdateDinosaur(c *gin.Context) {
	name := c.Param("name")

	var dinosaur models.Dinosaur
	if err := c.ShouldBindJSON(&dinosaur); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := dinosaur.UpdateDinosaur(name)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, dinosaur)
}

// DeleteDinosaur deletes a dinosaur by name
func (dc *DinosaurController) DeleteDinosaur(c *gin.Context) {
	name := c.Param("name")

	var dinosaur models.Dinosaur
	err := dinosaur.DeleteDinosaur(name)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Dinosaur deleted successfully!"})
}

func (dc *DinosaurController) FilterDinosaursBySpecies(c *gin.Context) {
	species := c.Query("species")
	if species == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	dinosaurs, err := models.GetDinosaursBySpecies(species)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, dinosaurs)
}

// Function to add a dinosaur to a cage
func (dc *DinosaurController) AddDinosaurToCage(c *gin.Context) {
	dinosaurName := c.Param("name")
	cageName := c.Param("cage_name")

	// Get the cage by name
	cage, err := models.GetCageByName(cageName)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	dinosaur, err := models.GetDinosaurByName(dinosaurName)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Add the dinosaur to the cage
	err = cage.AddDinosaur(&dinosaur)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dinosaur added to cage successfully!"})
}

// RemoveDinosaurFromCage removes the dinosaur with the given ID from its current cage
func (dc *DinosaurController) RemoveDinosaurFromCage(c *gin.Context) {
	dinosaurName := c.Param("name")

	dinosaur, err := models.GetDinosaurByName(dinosaurName)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Get the cage where the dinosaur is currently located
	cage, err := models.GetCageForDinosaur(dinosaurName)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Remove the dinosaur from its cage
	err = cage.RemoveDinosaur(&dinosaur)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Dinosaur removed from cage successfully!"})
}
