package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hereswilson/jurassic-park-api/controllers"
)

func speciesRoutes(superRoute *gin.RouterGroup) {
	speciesRouter := superRoute.Group("/species")
	{
		speciesRouter.GET("/", controllers.SpeciesControllers.GetSpecies)

		speciesRouter.POST("/", controllers.SpeciesControllers.CreateSpecies)

		speciesRouter.GET("/:id", controllers.SpeciesControllers.GetSpeciesByID)

		speciesRouter.PUT("/:id", controllers.SpeciesControllers.UpdateSpecies)

		speciesRouter.DELETE("/:id", controllers.SpeciesControllers.DeleteSpecies)
	}
}
