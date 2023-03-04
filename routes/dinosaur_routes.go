package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hereswilson/jurassic-park-api/controllers"
)

func dinosaurRoutes(superRoute *gin.RouterGroup) {
	dinosaurRouter := superRoute.Group("/dinosaurs")
	{
		dinosaurRouter.GET("/", controllers.GetDinosaurs)

		dinosaurRouter.POST("/", controllers.CreateDinosaur)

		dinosaurRouter.GET("/:name", controllers.GetDinosaurByName)

		dinosaurRouter.PUT("/:name", controllers.UpdateDinosaur)

		dinosaurRouter.DELETE("/:name", controllers.DeleteDinosaur)
	}
}
