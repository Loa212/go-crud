package main

import (
	"github/loa212/go-crud/initializers"
	"github/loa212/go-crud/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
