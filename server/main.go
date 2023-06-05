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

	
	// Define the routes
	router.POST("/shareproduct", controllers.ShareProduct)
    
	router.Run(":3000") // listen and serve on 0.0.0.0:8080
}
