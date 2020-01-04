package models

import (
	"mingi/goyoma/lib/common"

	"github.com/jinzhu/gorm"
)

// Content data model
type Content struct {
	gorm.Model
	Description string `sql:"type:text;"`
	User        User   `gorm:"foreignkey:UserID"`
	UserID      uint
	PoseList    []Pose
}

// Serialize serializes Content data
func (c Content) Serialize() common.JSON {
	return common.JSON{
		"id":          c.ID,
		"description": c.Description,
		"user":        c.User.Serialize(),
		"created_at":  c.CreatedAt,
	}
}
