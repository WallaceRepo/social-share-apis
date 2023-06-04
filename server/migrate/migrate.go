package main

import (
	"github.com/wallacerepo/social-share-apis/initializers"
	"github.com/wallacerepo/social-share-apis/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.Post{})
}

