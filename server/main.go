package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wallacerepo/social-share-apis/controllers"
	"github.com/wallacerepo/social-share-apis/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	fmt.Println("Hello World")
	router := gin.Default()

	router.POST("/posts", controllers.PostsCreate)

	router.PUT("/posts/:id", controllers.PostUpdate)

	router.DELETE("/posts/:id", controllers.PostDelete)

	router.GET("/posts", controllers.PostsGet)

	router.GET("/posts/:id", controllers.PostIndex)

	// share product routes
    
	router.POST("/share",)


	
	router.Run() // listen and serve on 0.0.0.0:8080
}
