package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/api/v1/user/:userName/login", UserLogin)
	router.GET("/api/v1/assets/:assetId/connect", WsConnect)

	router.Run(":8080")
}

func WsConnect(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	fmt.Println(token)
	u, ok := JwtVerify(token)
	if ok {

	}
	fmt.Println(u, ok)
}

func UserLogin(ctx *gin.Context) {
	token, err := JwtSign("admin")
	fmt.Println(token, err)
	ctx.JSON(200, token)
}
