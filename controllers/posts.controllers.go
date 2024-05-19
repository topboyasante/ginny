package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/topboyasante/ginny/internal/database"
	"github.com/topboyasante/ginny/internal/validators"
	"github.com/topboyasante/ginny/models"
)

func CreatePost(c *gin.Context) {
	var Body struct {
		Body  string
		Title string
	}

	// Parses the request body and stores it in the Body struct
	c.Bind(&Body)

	post := models.Post{Title: Body.Title, Body: Body.Body}

	// Validate the data from the request body
	if !validators.NotBlank(post.Title) || !validators.NotBlank(post.Body) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "some fields are empty",
		})
		return
	}

	res := database.DB.Create(&post)

	if res.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetAllPosts(c *gin.Context) {
	var posts []models.Post

	database.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPost(c *gin.Context) {
	var post models.Post

	// Get the id from the URL
	id := c.Param("id")

	database.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	var post models.Post

	var Body struct {
		Body  string
		Title string
	}

	// Parses the request body and stores it in the Body struct
	c.Bind(&Body)

	// Get the id from the URL
	id := c.Param("id")

	// Get the Post with the given ID
	database.DB.First(&post, id)

	database.DB.Model(&post).Updates(models.Post{
		Title: Body.Title,
		Body:  Body.Body,
	})

	c.JSON(200, gin.H{
		"post": Body,
	})
}

func DeletePost(c *gin.Context) {
	// Get the id from the URL
	id := c.Param("id")

	database.DB.Delete(&models.Post{}, id)

	c.JSON(200, gin.H{
		"success": "post deleted",
	})
}
