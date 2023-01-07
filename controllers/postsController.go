package controllers

import (
	"github/loa212/go-crud/initializers"
	"github/loa212/go-crud/models"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	//get data of req.body
	var body models.Post

	c.BindJSON(&body)
	
	//create post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Error creating post",
		})
		return
	}

	//return post
	c.JSON(200, gin.H{
		"message": "Post created",
		"post": post,
	})
	
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	
	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Error getting post with id=" + id,
		})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post

	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Error getting posts",
		})
		return
	}

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var body models.Post
	c.BindJSON(&body)

	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Error getting post with id=" + id,
		})
		return
	}

	post.Title = body.Title
	post.Body = body.Body

	result = initializers.DB.Save(&post)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Error updating post with id=" + id,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Post updated",
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Error getting post with id=" + id,
		})
		return
	}

	result = initializers.DB.Delete(&post)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Error deleting post with id=" + id,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Post deleted",
	})
}
