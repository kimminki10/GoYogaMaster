package contents

import (
	"mingi/goyoma/database/models"
	"mingi/goyoma/lib/common"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Content type alias
type Content = models.Content

// User type alias
type User = models.User

// JSON type alias
type JSON = common.JSON

func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		Description string `json:"description" building:"requiered"`
	}
	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}

	user := c.MustGet("user").(User)
	content := Content{Description: requestBody.Description, User: user}
	db.NewRecord(content)
	db.Create(&content)
	c.JSON(200, content.Serialize())
}

func list(c *gin.Context) {

}

func read(c *gin.Context) {

}

func remove(c *gin.Context) {

}

func update(c *gin.Context) {

}
