package main

import (
	"github.com/gin-gonic/gin"
	"github.com/topboyasante/ginny/controllers"
	"github.com/topboyasante/ginny/internal/config"
	"github.com/topboyasante/ginny/internal/database"
	"github.com/topboyasante/ginny/models"
)

// init() runs before the main function runs
func init() {
	database.ConnectToDB()
}

func main() {
	//Run AutoMigrations
	database.DB.AutoMigrate(&models.Post{})

	r := gin.Default()

	//Define your routes
	r.GET("/posts", controllers.GetAllPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.POST("/posts", controllers.CreatePost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	// listen and serve on 0.0.0.0:8080
	r.Run(config.ENV.ServerPort)
}
