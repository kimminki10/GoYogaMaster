package poses

import (
	"mingi/goyoma/database/models"
	"mingi/goyoma/lib/common"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

// Pose type alias
type Pose = models.Pose

// User type alias
type User = models.User

// JSON type alias
type JSON = common.JSON

func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		PoseName         string `json:"PoseName"`
		PoseDescription  string `json:"PoseDescription"`
		PoseThumbnailURL string `json:"PoseThumbnailURL"`
		PoseVideoURL     string `json:"PoseVideoURL"`
		PoseJSONURL      string `json:"PoseJSONURL"`
		PoseCategory     string `json:"PoseCategory"`
	}
	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}
	user := c.MustGet("user").(User)
	pose := Pose{
		PoseName:         requestBody.PoseName,
		PoseDescription:  requestBody.PoseDescription,
		PoseThumbnailURL: requestBody.PoseThumbnailURL,
		PoseVideoURL:     requestBody.PoseVideoURL,
		PoseJSONURL:      requestBody.PoseJSONURL,
		PoseCategory:     requestBody.PoseCategory,
		User:             user,
	}
	db.NewRecord(pose)
	db.Create(&pose)
	c.JSON(200, pose.Serialize())
}

func list(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	cursor := c.Query("cursor")
	recent := c.Query("recent")

	var poses []Pose

	if cursor == "" {
		if err := db.Preload("User").Limit(10).Order("id desc").Find(&poses).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	} else {
		condition := "id < ?"
		if recent == "1" {
			condition = "id > ?"
		}
		if err := db.Preload("User").Limit(10).Order("id desc").Where(condition, cursor).Find(&poses).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	}

	length := len(poses)
	serialized := make([]JSON, length, length)

	for i := 0; i < length; i++ {
		serialized[i] = poses[i].Serialize()
	}

	c.JSON(200, serialized)
}

func read(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var pose Pose

	// auto preloads the related model
	// http://gorm.io/docs/preload.html#Auto-Preloading
	if err := db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&pose).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, pose.Serialize())
}

func remove(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	user := c.MustGet("user").(User)

	var pose Pose
	if err := db.Where("id = ?", id).First(&pose).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	if pose.UserID != user.ID {
		c.AbortWithStatus(403)
		return
	}

	db.Delete(&pose)
	c.Status(204)
}

func update(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	user := c.MustGet("user").(User)

	type RequestBody struct {
		PoseName         string `json:"PoseName"`
		PoseDescription  string `json:"PoseDescription"`
		PoseThumbnailURL string `json:"PoseThumbnailURL"`
		PoseVideoURL     string `json:"PoseVideoURL"`
		PoseJSONURL      string `json:"PoseJSONURL"`
		PoseCategory     string `json:"PoseCategory"`
	}
	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}

	var pose Pose
	if err := db.Preload("User").Where("id = ?", id).First(&pose).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	if pose.UserID != user.ID {
		c.AbortWithStatus(403)
		return
	}

	pose.PoseName = requestBody.PoseName
	pose.PoseDescription = requestBody.PoseDescription
	pose.PoseThumbnailURL = requestBody.PoseThumbnailURL
	pose.PoseVideoURL = requestBody.PoseVideoURL
	pose.PoseJSONURL = requestBody.PoseJSONURL
	pose.PoseCategory = requestBody.PoseCategory
	db.Save(&pose)
	c.JSON(200, pose.Serialize())
}
