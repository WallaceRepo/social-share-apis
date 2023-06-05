package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func shareProduct(c *gin.Context, oauthConfig *oauth2.Config) {
	// Parse the request JSON
	var body struct {
		Product  string
		Title string
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Redirect the user to Facebook for authorization
	authURL := oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOnline)
	c.Redirect(http.StatusFound, authURL)
}
func handleCallback(c *gin.Context, oauthConfig *oauth2.Config) {
	// Retrieve the authorization code from the query parameters
	code := c.Query("code")

	// Exchange the authorization code for an access token
	token, err := oauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange authorization code"})
		return
	}

	// Get the access token from the token response
	accessToken := token.AccessToken

	// Create the post message
	// post := fmt.Sprintf("Check out this amazing product: %s", product.Name)
	post := fmt.Sprintf("Check out this amazing product" )


	// Share the post
	err = shareToFacebook(accessToken, post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to share product"})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "Product shared successfully"})
}

func shareToFacebook(accessToken string, post string) error {
	url := fmt.Sprintf("https://graph.facebook.com/me/feed?access_token=%s&message=%s", accessToken, post)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to share post. Status: %s", resp.Status)
	}

	return nil
}
