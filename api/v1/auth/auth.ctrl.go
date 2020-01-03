package auth

import (
	"io/ioutil"
	"mingi/goyoma/database/models"
	"mingi/goyoma/lib/common"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User = models.User

func hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func checkHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateToken(data common.JSON) (string, error) {

	//  token is valid for 7days
	date := time.Now().Add(time.Hour * 24 * 7)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": data,
		"exp":  date.Unix(),
	})

	// get path from root dir
	pwd, _ := os.Getwd()
	keyPath := pwd + "/jwtsecret.key"

	key, readErr := ioutil.ReadFile(keyPath)
	if readErr != nil {
		return "", readErr
	}
	tokenString, err := token.SignedString(key)
	return tokenString, err
}

func register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		Email    string  `json:"email" binding:"required"`
		Password string  `json:"password" binding:"required"`
		UFName   string  `json:"UFName" binding:"required"`
		ULName   string  `json:"ULName" binding:"required"`
		Mobile   string  `json:"Mobile" binding:"required"`
		Height   float32 `json:"Height" binding:"required"`
		Weight   float32 `json:"Weight" binding:"required"`
		Age      int     `json:"Age" binding:"required"`
		Gender   string  `json:"Gender" binding:"required"`
		UType    string  `json:"UType" binding:"required"`
	}

	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatus(400)
		return
	}

	// check existancy
	var exists User
	if err := db.Where("email = ?", body.Email).First(&exists).Error; err == nil {
		c.AbortWithStatus(409)
		return
	}

	hash, hashErr := hash(body.Password)
	if hashErr != nil {
		c.AbortWithStatus(500)
		return
	}

	// create user
	user := User{
		Email:        body.Email,
		UFName:       body.UFName,
		ULName:       body.ULName,
		Mobile:       body.Mobile,
		Height:       body.Height,
		Weight:       body.Weight,
		Age:          body.Age,
		Gender:       body.Gender,
		UType:        body.UType,
		PasswordHash: hash,
	}

	db.NewRecord(user)
	db.Create(&user)

	serialized := user.Serialize()
	token, _ := generateToken(serialized)
	c.SetCookie("token", token, 60*60*24*7, "/", "", false, true)

	c.JSON(200, common.JSON{
		"user":  user.Serialize(),
		"token": token,
	})
}

func login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatus(400)
		return
	}

	// check existancy
	var user User
	if err := db.Where("email = ?", body.Email).First(&user).Error; err != nil {
		c.AbortWithStatus(404) // user not found
		return
	}

	if !checkHash(body.Password, user.PasswordHash) {
		c.AbortWithStatus(401)
		return
	}

	serialized := user.Serialize()
	token, _ := generateToken(serialized)

	c.SetCookie("token", token, 60*60*24*7, "/", "", false, true)

	c.JSON(200, common.JSON{
		"user":  user.Serialize(),
		"token": token,
	})

}
