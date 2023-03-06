package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/hereswilson/jurassic-park-api/controllers"
	"github.com/hereswilson/jurassic-park-api/services"
	"github.com/hereswilson/jurrasic-park-api/repositories"

	"gorm.io/gorm"
)

func cageRoutes(superRoute *gin.RouterGroup, db *gorm.DB) {
	cageRepo := repositories.NewCageRepository(db)
	cageService := services.NewCageService(cageRepo)
	cageController := controllers.NewCageController(cageService)
	cageRouter := superRoute.Group("/cages")
	{
		cageRouter.GET("/", cageController.GetCages)

		cageRouter.POST("/", cageController.CreateCage)

		cageRouter.GET("/:name", cageController.GetCageByName)

		cageRouter.PUT("/:name", cageController.UpdateCage)

		cageRouter.DELETE("/:name", cageController.DeleteCage)

		cageRouter.GET("/:name/dinosaurs", cageController.GetDinosaursInCage)
	}
}
