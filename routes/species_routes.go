package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hereswilson/jurassic-park-api/controllers"
)

func speciesRoutes(superRoute *gin.RouterGroup) {
	speciesController := &controllers.SpeciesController{}
	speciesRouter := superRoute.Group("/species")
	{
		speciesRouter.GET("/", speciesController.GetSpecies)

		speciesRouter.POST("/", speciesController.CreateSpecies)

		speciesRouter.GET("/:species", speciesController.GetSpecificSpecies)

		speciesRouter.PUT("/:species", speciesController.UpdateSpecies)

		speciesRouter.DELETE("/:species", speciesController.DeleteSpecies)
	}
}
