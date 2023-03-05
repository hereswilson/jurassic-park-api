package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hereswilson/jurassic-park-api/controllers"
)

func cageRoutes(superRoute *gin.RouterGroup) {
	cageController := &controllers.CageController{}
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
