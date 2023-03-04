package routes

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(superRoute *gin.RouterGroup) {
	cageRoutes(superRoute)
	dinosaurRoutes(superRoute)
	speciesRoutes(superRoute)
}
