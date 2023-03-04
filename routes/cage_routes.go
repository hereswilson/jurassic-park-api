package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hereswilson/jurassic-park-api/controllers"
)

func cageRoutes(superRoute *gin.RouterGroup) {
	cageRouter := superRoute.Group("/cages")
	{
		cageRouter.GET("/", controllers.CageControllers.GetCages)

		cageRouter.POST("/", controllers.CageControllers.CreateCage)

		cageRouter.GET("/:id", controllers.CageControllers.GetCageByID)

		cageRouter.PUT("/:id", controllers.CageControllers.UpdateCage)

		cageRouter.DELETE("/:id", controllers.CageControllers.DeleteCage)
	}
}
