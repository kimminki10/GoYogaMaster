package main

import (
	"mingi/goyoma/api"
	"mingi/goyoma/database"
	"mingi/goyoma/lib/middlewares"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, _ := database.Initialize()
	port := os.Getenv("PORT")
	app := gin.Default()
	app.Use(database.Inject(db))
	app.Use(middlewares.JWTMiddleware())
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	app.Use(sessions.Sessions("redissession", store))
	api.ApplyRoutes(app)
	app.Run(":" + port)
}
