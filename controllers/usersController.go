package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoannguyen02/self-go/models"
	"github.com/hoannguyen02/self-go/services"
)

func SignUp(c *gin.Context) {
	user := new(models.User)
	err := c.Bind(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	}
	// Hash pwd
	// Create user
	err = services.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}

func SignIn(c *gin.Context) {
	// Bind user request body
	var user models.User
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	}
	// Get user by username
	user, err = services.GetUserByUsername(user.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username or password is invalid",
		})
		return
	}
	// Validate password
	// Generate session
	// Save session as a key with user value to redis 
	// Save session id to cookie
	// c.SetSameSite(http.SameSiteLaxMode)
	// c.SetCookie("Authorization", "sessionId", 3600*24*30, "", "", false, true)
	// Response
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})

}