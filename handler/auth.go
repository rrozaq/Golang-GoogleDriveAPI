package handler

import (
	"errors"
	"golang.org/x/oauth2"

	"gdrive_uploader/pkg/auth"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	auth, err := auth.Authenticate()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "authentication failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "authentication success",
		"auth":    auth,
	})
}

func Callback(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")

	err := getUserInfo(code, state)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func getUserInfo(code string, state string) error {
	config, err := auth.ClientConfig()
	if err != nil {
		return err
	}

	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return errors.New("failed to exchange token")
	}

	err = auth.WriteTokenFile(token)
	if err != nil {
		return err
	}

	return nil
}
