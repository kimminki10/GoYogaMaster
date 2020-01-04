package poses

import (
	"mingi/goyoma/lib/middlewares"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	poses := r.Group("/poses")
	{
		poses.POST("/", middlewares.Authorized, create)
		poses.GET("/", list)
		poses.GET("/:id", read)
		poses.DELETE("/:id", middlewares.Authorized, remove)
		poses.PATCH("/:id", middlewares.Authorized, update)
	}
}
