package models

import (
	"mingi/goyoma/lib/common"

	"github.com/jinzhu/gorm"
)

// Pose data model
type Pose struct {
	gorm.Model
	PoseName         string `json:"pose_name"`
	PoseDescription  string
	PoseThumbnailURL string
	PoseVideoURL     string
	PoseJSONURL      string
	PoseCategory     string
	User             User
	UserID           uint `gorm:"foreignkey:UserID"`
}

// Serialize serializes pose data
func (p Pose) Serialize() common.JSON {
	return common.JSON{
		"id":               p.ID,
		"PoseName":         p.PoseName,
		"PoseDescription":  p.PoseDescription,
		"PoseThumbnailURL": p.PoseThumbnailURL,
		"PoseVideoURL":     p.PoseVideoURL,
		"PoseJSONURL":      p.PoseJSONURL,
		"PoseCategory":     p.PoseCategory,
		"User":             p.User.Serialize(),
		"created_at":       p.CreatedAt,
	}
}
