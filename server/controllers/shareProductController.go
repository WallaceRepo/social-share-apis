package controllers

import (
	// "net/http"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wallacerepo/social-share-apis/initializers"
	"github.com/wallacerepo/social-share-apis/models"
	"github.com/emrearmagan/go-social"
)

// func ShareProduct(ctx *gin.Context) {

// 	// Check if user is signed in
// 	if !c.isUserSignedIn(ctx) {
// 		ctx.Redirect(http.StatusFound, "/signin")
// 		return
// 	}
//      // chosen platform by user
// 	platform := ctx.PostForm("platform")

// }

//	func (c *UserController) isUserSignedIn(ctx *gin.Context) bool {
//		// Check if the user is signed in
//		// This could involve checking for a valid session or token
//		// Example implementation for session-based authentication:
//		session := sessions.Default(ctx)
//		userId := session.Get("userId")
//		return userId != nil
//	}
type ShareProductController struct {
	SocialAuth *social.SocialAuth
}

func (ctrl *ShareProductController) ShareProductCallback(c *gin.Context) {
	// Parse the callback response from the social media platform
	providerName := c.Query("provider")
	code := c.Query("code")

	// Get the provider based on the provider name
	provider, err := ctrl.SocialAuth.GetProviderByName(providerName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid social media platform"})
		return
	}

	// Exchange the authorization code for an access token
	err = provider.Exchange(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange authorization code"})
		return
	}

	// Get the shared product from the database
	var sharedProduct models.SharedProduct
	err = initializers.DB.Last(&sharedProduct).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Shared product not found"})
		return
	}

	// Save the user media token in the database
	userMediaToken := models.UserMediaToken{
		Username:     provider.GetUsername(),
		AccessToken:  provider.GetAccessToken(),
		TokenType:    provider.GetTokenType(),
		ExpiresIn:    provider.GetExpiresIn(),
		RefreshToken: provider.GetRefreshToken(),
	}
	if err := initializers.DB.Create(&userMediaToken).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user media token"})
		return
	}

	// Update the shared product with the user media token
	sharedProduct.UserID = userMediaToken.Username
	sharedProduct.Link = "" // Set the link based on the social media platform's API response
	if err := initializers.DB.Save(&sharedProduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update shared product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product shared successfully"})
}
