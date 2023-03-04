package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hereswilson/jurassic-park-api/controllers"
)

func cageRoutes(superRoute *gin.RouterGroup) {
	cageRouter := superRoute.Group("/cages")
	{
		cageRouter.GET("/", controllers.GetCages)

		cageRouter.POST("/", controllers.CreateCage)

		cageRouter.GET("/:name", controllers.GetCageByName)

		cageRouter.PUT("/:name", controllers.UpdateCage)

		cageRouter.DELETE("/:name", controllers.DeleteCage)
	}
}
