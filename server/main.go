package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wallacerepo/social-share-apis/controllers"
	"github.com/wallacerepo/social-share-apis/initializers"
	"golang.org/x/oauth2"
	// "golang.org/x/oauth2/facebook"
)

var (
	oauthConfig *oauth2.Config
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	//_, APP_ID, APP_SECRET, REDIRECT_URL := initializers.LoadEnvVariables()

	// Initialize the OAuth2 configuration
	// oauthConfig = &oauth2.Config {
	// 	ClientID:     APP_ID,
	// 	ClientSecret: APP_SECRET,
	// 	RedirectURL:  REDIRECT_URL,
	// 	Scopes:       []string{"publish_actions"},
	// 	Endpoint: facebook.Endpoint,
	// }
}

func main() {
	fmt.Println("Hello World")
	router := gin.Default()

    router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
	router.POST("/posts", controllers.PostsCreate)

	router.PUT("/posts/:id", controllers.PostUpdate)

	router.DELETE("/posts/:id", controllers.PostDelete)

	router.GET("/posts", controllers.PostsGet)

	router.GET("/posts/:id", controllers.PostIndex)

	// route for sharing or posting product to fbook

	router.POST("/share", controllers.ShareProduct)

	// Define the callback route for Facebook authorization
	router.GET("/callback", controllers.HandleCallback)

	router.Run(":3000") // listen and serve on 0.0.0.0:8080
}
