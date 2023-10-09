package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)	

func getSessionId() (string, error) {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}

type User struct {
	Username string
	Password string
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB: 0,
		Password: "",
	})
	// print(client)
	router := gin.Default()
	api := router.Group("/api") 
	{
		api.GET("/ping", func(ctx *gin.Context) {
			// client.Incr(ctx)
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		api.POST("/login", func(ctx *gin.Context) {
			// Decode the request body into the User struct
			var user User
			err := json.NewDecoder(ctx.Request.Body).Decode(&user)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, "")
			}
			// Generate session id
			sessionId, err := getSessionId()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, "")
			}
			// Set session id as a key with username value
			err = client.Set(sessionId, user.Username, 0).Err()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, "")
			}
			// Get username with session id key
			val, err := client.Get(sessionId).Result()
			if err != nil {
				fmt.Println(err)
			}
			// Send success http response
			ctx.JSON(http.StatusOK, gin.H{
				"username": val,
			})
		})
	}

	router.Run()


}