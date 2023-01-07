package main

import (
	"github/loa212/go-crud/controllers"
	"github/loa212/go-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	r.POST("/post", controllers.CreatePost)
	r.GET("/post/:id", controllers.GetPost)
	r.PUT("/post/:id", controllers.UpdatePost)
	r.DELETE("/post/:id", controllers.DeletePost)
	r.GET("/getPosts", controllers.GetPosts)
	r.Run() // listen and serve on 0.0.0.0:3000 (for windows "localhost:3000")
}