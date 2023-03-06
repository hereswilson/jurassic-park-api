package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hereswilson/jurassic-park-api/controllers"
	"github.com/hereswilson/jurassic-park-api/services"
	"github.com/hereswilson/jurrasic-park-api/repositories"
	"gorm.io/gorm"
)

func speciesRoutes(superRoute *gin.RouterGroup, db *gorm.DB) {
	speciesRepo := repositories.NewSpeciesRepository(db)
	speciesService := services.NewSpeciesService(speciesRepo)
	speciesController := controllers.NewSpeciesController(speciesService)
	speciesRouter := superRoute.Group("/species")
	{
		speciesRouter.GET("/", speciesController.GetSpecies)

		speciesRouter.POST("/", speciesController.CreateSpecies)

		speciesRouter.GET("/:species", speciesController.GetSpecificSpecies)

		speciesRouter.PUT("/:species", speciesController.UpdateSpecies)

		speciesRouter.DELETE("/:species", speciesController.DeleteSpecies)
	}
}
