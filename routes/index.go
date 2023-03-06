package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddRoutes(superRoute *gin.RouterGroup, db *gorm.DB) {
	cageRoutes(superRoute, db)
	dinosaurRoutes(superRoute, db)
	speciesRoutes(superRoute, db)
}
