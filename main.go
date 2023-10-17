/*
1 API /login, để tạo session cho mỗi người đăng nhập, dùng redis để lưu session id, user name ấy

1 API /ping chỉ cho phép 1 người được gọi tại một thời điểm ( với sleep ở bên trong api đó trong 5s)

đếm số lượng lần 1 người gọi api /ping

rate limit mỗi người chỉ được gọi API /ping 2 lần trong 60s

1 API /top/ trả về top 10 người gọi API /ping nhiều nhất

Dùng hyperloglog để lưu xấp sỉ số người gọi api /ping , và trả về trong api /count
*/

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoannguyen02/self-go/controllers"
	"github.com/hoannguyen02/self-go/redis"
)

func init() {
	redis.CreateClient()
}

func main() {
	run();
}

type Session struct {
	Id string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func run() {
	router := gin.Default()
	api := router.Group("/api") 
	{
		api.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		api.POST("/signup", controllers.SignUp)
		api.POST("/login", controllers.SignIn)
	}

	router.Run()

}