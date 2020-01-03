package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/velopert/gin-rest-api-sample/api/v1.0/auth"
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
	}
}
