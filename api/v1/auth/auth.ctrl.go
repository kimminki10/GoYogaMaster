package auth

import (
	"mingi/goyoma/database/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type User = models.User

func register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		email string `json:"email" binding:"required"`
		password
	}

	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatus(400)
		return
	}

	// check existancy
	var exists User
	if err := db.Where("email = ?", body.email).First(&exists).Error; err == nil {
		c.AbortWithStatus(409)
		return
	}

	hash, hashErr := hash(body.Password)
}
