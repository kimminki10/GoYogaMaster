package models

import (
	"mingi/goyoma/lib/common"

	"github.com/jinzhu/gorm"
)

// Content data model
type Content struct {
	gorm.Model
	ContentName         string
	Description         string `sql:"type:text;"`
	Durations           float64
	Category            string
	ContentThumbnailURL string
	ContentViewNum      uint
	ContentStatus       string
	User                User
	UserID              uint `gorm:"foreignkey:UserID"`
	PoseList            []Pose
}

// Serialize serializes Content data
func (c Content) Serialize() common.JSON {
	return common.JSON{
		"id":                  c.ID,
		"description":         c.Description,
		"user":                c.User.Serialize(),
		"created_at":          c.CreatedAt,
		"ContentName":         c.ContentName,
		"Durations":           c.Durations,
		"category":            c.Category,
		"ContentThumbnailURL": c.ContentThumbnailURL,
		"ContentViewNum":      c.ContentViewNum,
		"ContentStatus":       c.ContentStatus,
		"PoseList":            c.PoseList,
	}
}
