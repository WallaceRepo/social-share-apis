package main

import (

	// "github.com/emrearmagan/go-social/social/twitter"

	//"github.com/mrjones/oauth"

	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dghubble/oauth1"

	"github.com/gin-gonic/gin"
	"github.com/wallacerepo/social-share-apis/controllers"
	"github.com/wallacerepo/social-share-apis/initializers"
	// "github.com/emrearmagan/go-social/oauth"
	// "log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	_, _, _, CONSUMER_SECRET, CONSUMER_KEY := initializers.LoadEnvVariables()

	fields := "created_at,description"
	params := map[string]string{
		"user.fields": fields,
	}

	// Get request token
	//	requestTokenURL := "https://api.twitter.com/oauth/request_token"
	//	accessTokenURL := "https://api.twitter.com/oauth/access_token"

	authorizationURL := "https://api.twitter.com/oauth/authorize"
	oauthConfig := oauth1.NewConfig(CONSUMER_KEY, CONSUMER_SECRET)
	oauthToken := oauth1.NewToken("", "")

	// Get authorization URL
	authorizationURL, _, err := oauthConfig.AuthorizationURL(authorizationURL)
	if err != nil {
		fmt.Printf("Failed to obtain authorization URL: %v\n", err)
		return
	}
	authorizationURLString := authorizationURL.String()

	// Wait for user authorization
	fmt.Print("Paste the PIN here: ")
	var verifier string
	_, err = fmt.Scan(&verifier)
	if err != nil {
		fmt.Printf("Failed to read verifier: %v\n", err)
		return
	}

	// Get access token
	accessToken, accessSecret, err := oauthConfig.AccessToken(nil, oauthToken, verifier)
	if err != nil {
		fmt.Printf("Failed to obtain access token: %v\n", err)
		return
	}

	// Create the OAuth1 client
	config := oauth1.NewConfig(CONSUMER_KEY, CONSUMER_SECRET)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(nil, token)

	// Make the request
	resourceURL := "https://api.twitter.com/2/users/me"
	resp, err := httpClient.Get(resourceURL + "?user.fields=" + fields)
	if err != nil {
		fmt.Printf("Failed to make the request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Request returned an error: %s\n", resp.Status)
		return
	}

	// Read response body
	var jsonResponse interface{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResponse)
	if err != nil {
		fmt.Printf("Failed to parse response: %v\n", err)
		return
	}

	// Print JSON response
	prettyJSON, err := json.MarshalIndent(jsonResponse, "", "    ")
	if err != nil {
		fmt.Printf("Failed to format JSON response: %v\n", err)
		return
	}
	fmt.Println(string(prettyJSON))

	// Start Gin server
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, jsonResponse)
	})

	// Define the routes
	router.POST("/shareproduct", controllers.ShareProduct)

	router.Run(":3000") // listen and serve on 0.0.0.0:8080
}
