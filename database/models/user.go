package models

import (
	"mingi/goyoma/lib/common"

	"github.com/jinzhu/gorm"
)

// User Data model
type User struct {
	gorm.Model
	Email        string
	PasswordHash string
	UFName       string
	ULName       string
	Mobile       string
	Height       float32
	Weight       float32
	Age          int
	Gender       string

	UType string
}

// Serialize serializes user data
func (u *User) Serialize() common.JSON {
	result := common.JSON{
		"id":            u.ID,
		"email":         u.Email,
		"password_hash": u.PasswordHash,
		"u_f_name":      u.UFName,
		"u_l_name":      u.ULName,
		"mobile":        u.Mobile,
		"height":        u.Height,
		"weight":        u.Weight,
		"age":           u.Age,
		"gender":        u.Gender,
		"u_type":        u.UType,
	}
	return result
}

func (u *User) Read(m common.JSON) {
	u.Email = m["email"].(string)
	u.PasswordHash = m["password_hash"].(string)
	u.UFName = m["u_f_name"].(string)
	u.ULName = m["u_l_name"].(string)
	u.Mobile = m["mobile"].(string)
	u.Height = m["height"].(float32)
	u.Weight = m["weight"].(float32)
	u.Age = m["age"].(int)
	u.Gender = m["gender"].(string)
	u.UType = m["u_type"].(string)
}
