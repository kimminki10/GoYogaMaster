package contents

import (
	"mingi/goyoma/lib/middlewares"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	contents := r.Group("/contents")
	{
		contents.POST("/", middlewares.Authorized, create)
		contents.GET("/", list)
		contents.GET("/:id", read)
		contents.DELETE("/:id", middlewares.Authorized, remove)
		contents.PATCH("/:id", middlewares.Authorized, update)
	}
}
