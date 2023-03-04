package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hereswilson/jurassic-park-api/controllers"
)

func dinosaurRoutes(superRoute *gin.RouterGroup) {
	dinosaurRouter := superRoute.Group("/dinosaurs")
	{
		dinosaurRouter.GET("/", controllers.DinosaurControllers.GetDinosaurs)

		dinosaurRouter.POST("/", controllers.DinosaurControllers.CreateDinosaur)

		dinosaurRouter.GET("/:id", controllers.DinosaurControllers.GetDinosaurByID)

		dinosaurRouter.PUT("/:id", controllers.DinosaurControllers.UpdateDinosaur)

		dinosaurRouter.DELETE("/:id", controllers.DinosaurControllers.DeleteDinosaur)
	}
}
