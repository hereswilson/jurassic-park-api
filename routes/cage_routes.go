package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/hereswilson/jurassic-park-api/controllers"
	"github.com/hereswilson/jurassic-park-api/repositories"
	"github.com/hereswilson/jurassic-park-api/services"

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

		cageRouter.GET("/name", cageController.GetCageByName)

		cageRouter.PUT("/", cageController.UpdateCage)

		cageRouter.DELETE("/", cageController.DeleteCage)

		cageRouter.GET("/dinosaurs", cageController.GetDinosaursInCage)
	}
}
