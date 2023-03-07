package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hereswilson/jurassic-park-api/controllers"
	"github.com/hereswilson/jurassic-park-api/repositories"
	"github.com/hereswilson/jurassic-park-api/services"
	"gorm.io/gorm"
)

func dinosaurRoutes(superRoute *gin.RouterGroup, db *gorm.DB) {
	dinoRepo := repositories.NewDinosaurRepository(db)
	cageRepo := repositories.NewCageRepository(db)
	speciesRepo := repositories.NewSpeciesRepository(db)
	dinoService := services.NewDinosaurService(dinoRepo, cageRepo, speciesRepo)
	dinosaurController := controllers.NewDinosaurController(dinoService)
	dinosaurRouter := superRoute.Group("/dinosaurs")
	{
		dinosaurRouter.GET("/", dinosaurController.GetDinosaurs)

		dinosaurRouter.POST("/", dinosaurController.CreateDinosaur)

		dinosaurRouter.GET("/name", dinosaurController.GetDinosaurByName)

		dinosaurRouter.PUT("/", dinosaurController.UpdateDinosaur)

		dinosaurRouter.DELETE("/", dinosaurController.DeleteDinosaur)

		dinosaurRouter.POST("/add-to-cage/", dinosaurController.AddDinosaurToCage)

		dinosaurRouter.POST("/remove-from-cage/", dinosaurController.RemoveDinosaurFromCage)

		dinosaurRouter.GET("/filter-by-species/", dinosaurController.FilterDinosaursBySpecies)
	}
}
