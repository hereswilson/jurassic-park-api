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

		dinosaurRouter.GET("/", dinosaurController.GetDinosaurByName)

		dinosaurRouter.PUT("/:name", dinosaurController.UpdateDinosaur)

		dinosaurRouter.DELETE("/:name", dinosaurController.DeleteDinosaur)

		dinosaurRouter.POST("/:name/add-to-cage/:cage_name", dinosaurController.AddDinosaurToCage)

		dinosaurRouter.POST("/:name/remove-from-cage", dinosaurController.RemoveDinosaurFromCage)

		dinosaurRouter.GET("/filter-by-species/:species", dinosaurController.FilterDinosaursBySpecies)
	}
}
