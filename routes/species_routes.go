package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hereswilson/jurassic-park-api/controllers"
)

func speciesRoutes(superRoute *gin.RouterGroup) {
	speciesRouter := superRoute.Group("/species")
	{
		speciesRouter.GET("/", controllers.GetSpecies)

		speciesRouter.POST("/", controllers.CreateSpecies)

		speciesRouter.GET("/:species", controllers.GetSpecificSpecies)

		speciesRouter.PUT("/:species", controllers.UpdateSpecies)

		speciesRouter.DELETE("/:species", controllers.DeleteSpecies)
	}
}
