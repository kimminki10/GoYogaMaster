package apiv1

import (
	"mingi/goyoma/api/v1/auth"
	"mingi/goyoma/api/v1/contents"
	"mingi/goyoma/api/v1/poses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		v1.GET("/ping", ping)
		auth.ApplyRoutes(v1)
		poses.ApplyRoutes(v1)
		contents.ApplyRoutes(v1)
	}
}
