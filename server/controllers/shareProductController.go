package controllers

import (
	"net/http"
	"github.com/emrearmagan/go-social"
	
	"github.com/gin-gonic/gin"
	"github.com/wallacerepo/social-share-apis/initializers"
	"github.com/wallacerepo/social-share-apis/models"
)

func ShareProduct(c *gin.Context) {

	// Get data off req body
	var body struct {
		UserID      int    `json:"userId"`
		ProductID   int    `json:"productId"`
		SocialMedia string `json:"socialMedia"`
		Link        string `gorm:"not null"`
	}

	c.Bind(&body)

	sharedProd := models.SharedProduct{
		UserID:      body.UserID, // for testing perpose
		ProductID:   body.ProductID,
		SocialMedia: body.SocialMedia,
	}

	result := initializers.DB.Create(&sharedProd)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Retrieve the logged-in user from the session or database
	// user := getCurrentUser(c)
	user := body.UserID

	// Get the selected social media platform from the request
	// socialMedia := c.PostForm("social_media")

	socialMedia := body.SocialMedia

	// Initialize the go-social library with the appropriate social media platform
	social := social.New(socialMedia)

	// Redirect the user to the social media platform for authentication
	authURL := social.AuthURL("YOUR_REDIRECT_URL")
	c.Redirect(http.StatusTemporaryRedirect, authURL)
}

func getCurrentUser(c *gin.Context) *models.User {
	// Retrieve the user from the session or database based on your authentication mechanism

	// Return the current user object
}
