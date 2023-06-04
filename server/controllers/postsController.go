package controllers

import (
	// "net/http"

	"github.com/gin-gonic/gin"
	"github.com/wallacerepo/social-share-apis/initializers"
	"github.com/wallacerepo/social-share-apis/models"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	// Create a post

	// post := models.Post{Title: "Hello", Body: "Post Body"}

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsGet(c *gin.Context) {
	// Retrieve all posts from the database

	// Get all posts
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return the posts
	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostIndex(c *gin.Context) {
	// Retrieve single posts from the database

	// Get id off url
	id := c.Param("id")

	// Get the post
	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return the post
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostUpdate(c *gin.Context) {
	// Get the id off the url
       id := c.Param("id")

	// Get the data off req body

	var body struct {
		Body string
		Title string
	}
	c.Bind(&body)

	// Find the post where updated
     
	var post models.Post
	initializers.DB.First(&post, id)


	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context){
    // Get the id off the url
     id := c.Param("id")

	// Delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
    
	c.Status(200)

}