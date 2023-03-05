package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hereswilson/jurassic-park-api/controllers"
)

func dinosaurRoutes(superRoute *gin.RouterGroup) {
	dinosaurController := &controllers.DinosaurController{}
	dinosaurRouter := superRoute.Group("/dinosaurs")
	{
		dinosaurRouter.GET("/", dinosaurController.GetDinosaurs)

		dinosaurRouter.POST("/", dinosaurController.CreateDinosaur)

		dinosaurRouter.GET("/:name", dinosaurController.GetDinosaurByName)

		dinosaurRouter.PUT("/:name", dinosaurController.UpdateDinosaur)

		dinosaurRouter.DELETE("/:name", dinosaurController.DeleteDinosaur)

		dinosaurRouter.POST("/:name/add-to-cage/:cage_name", dinosaurController.AddDinosaurToCage)

		dinosaurRouter.POST("/:name/remove-from-cage", dinosaurController.RemoveDinosaurFromCage)

		dinosaurRouter.GET("/filter-by-species/:species", dinosaurController.FilterDinosaursBySpecies)
	}
}
