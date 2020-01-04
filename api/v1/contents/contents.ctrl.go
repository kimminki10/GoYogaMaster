package contents

import (
	"fmt"
	"mingi/goyoma/database/models"
	"mingi/goyoma/lib/common"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Content type alias
type Content = models.Content

// User type alias
type User = models.User

// Pose type alias
type Pose = models.Pose

// JSON type alias
type JSON = common.JSON

func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		ContentName         string  `json:"ContentName" building:"requiered"`
		Description         string  `json:"Description"`
		Durations           float64 `json:"Durations"`
		Category            string  `json:"Category"`
		ContentThumbnailURL string  `json:"ContentThumbnailURL"`
		ContentViewNum      uint    `json:"ContentViewNum"`
		ContentStatus       string  `json:"ContentStatus"`
		PoseList            []uint  `json:"PoseList"`
	}
	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}

	var poseList []Pose
	condition := "id IN (?)"
	if err := db.Set("gorm:auto_preload", true).Where(condition, requestBody.PoseList).Find(&poseList).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatus(404)
		return
	}

	user := c.MustGet("user").(User)

	content := Content{
		ContentName:         requestBody.ContentName,
		Description:         requestBody.Description,
		Durations:           requestBody.Durations,
		Category:            requestBody.Category,
		ContentThumbnailURL: requestBody.ContentThumbnailURL,
		ContentViewNum:      requestBody.ContentViewNum,
		ContentStatus:       requestBody.ContentStatus,
		PoseList:            poseList,
		User:                user,
	}
	db.NewRecord(content)
	db.Create(&content)
	c.JSON(200, content.Serialize())
}

func list(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	cursor := c.Query("cursor")
	recent := c.Query("recent")

	var contents []Content

	if cursor == "" {
		if err := db.Preload("User").Limit(10).Order("id desc").Find(&contents).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	} else {
		condition := "id < ?"
		if recent == "1" {
			condition = "id > ?"
		}
		if err := db.Preload("User").Limit(10).Order("id desc").Where(condition, cursor).Find(&contents).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	}

	length := len(contents)
	serialized := make([]JSON, length, length)

	for i := 0; i < length; i++ {
		serialized[i] = contents[i].Serialize()
	}

	c.JSON(200, serialized)
}

func read(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var content Content

	// auto preloads the related model
	// http://gorm.io/docs/preload.html#Auto-Preloading
	if err := db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&content).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, content.Serialize())
}

func remove(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	user := c.MustGet("user").(User)

	var content Content
	if err := db.Where("id = ?", id).First(&content).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	if content.UserID != user.ID {
		c.AbortWithStatus(403)
		return
	}

	db.Delete(&content)
	c.Status(204)
}

func update(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	user := c.MustGet("user").(User)

	type RequestBody struct {
		ContentName         string  `json:"ContentName"`
		Description         string  `json:"Description"`
		Durations           float64 `json:"Durations"`
		Category            string  `json:"Category"`
		ContentThumbnailURL string  `json:"ContentThumbnailURL"`
		ContentViewNum      uint    `json:"ContentViewNum"`
		ContentStatus       string  `json:"ContentStatus"`
	}

	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}

	var content Content
	if err := db.Preload("User").Where("id = ?", id).First(&content).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	if content.UserID != user.ID {
		c.AbortWithStatus(403)
		return
	}

	content.ContentName = requestBody.ContentName
	content.Description = requestBody.Description
	content.Durations = requestBody.Durations
	content.Category = requestBody.Category
	content.ContentThumbnailURL = requestBody.ContentThumbnailURL
	content.ContentViewNum = requestBody.ContentViewNum
	content.ContentStatus = requestBody.ContentStatus
	db.Save(&content)
	c.JSON(200, content.Serialize())
}
