package middlewares

import (
	"mingi/goyoma/database/models"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Authorized check user is exist in context, and check token is in redis session
func Authorized(c *gin.Context) {
	userRaw, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(401)
		return
	}
	session := sessions.Default(c)

	user := userRaw.(models.User)
	redisToken := session.Get(user.Email)
	tokenString, err := c.Cookie("token")
	if err != nil {
		// try reading HTTP Header
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			c.AbortWithStatus(401)
			return
		}
		sp := strings.Split(authorization, "Bearer ")
		// invalid token
		if len(sp) < 1 {
			c.AbortWithStatus(401)
			return
		}
		tokenString = sp[1]
	}

	if redisToken != tokenString {
		c.AbortWithStatus(401)
		return
	}
}
