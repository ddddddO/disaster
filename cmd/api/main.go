package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("api server started")

	router := gin.Default()

	// TODO: 切り出す
	// https://github.com/gin-gonic/gin#grouping-routes
	a := router.Group("/weather")
	a.GET("", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "%s %s\n", "GET", "weather")
	})

	b := router.Group("/earthquake")
	b.GET("", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "%s %s\n", "GET", "earthquake")
	})

	c := router.Group("/blackout")
	c.GET("", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "%s %s\n", "GET", "blackout")
	})

	d := router.Group("/basestation")
	d.GET("", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "%s %s\n", "GET", "basestation")
	})

	e := router.Group("/asylum")
	e.GET("", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "%s %s\n", "GET", "asylum")
	})

	f := router.Group("/localthreads")
	f.GET("", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "%s %s\n", "GET", "localthreads")
	})

	g := router.Group("/chat")
	g.GET("", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "%s %s\n", "GET", "chat")
	})
	g.POST("/user", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "%s %s\n", "POST", "chat/user")
	})
	g.POST("/login", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "%s %s\n", "POST", "chat/login")
	})
	g.POST("/channels", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "%s %s\n", "POST", "chat/channels")
	})
	g.GET("/channels", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "%s %s\n", "GET", "chat/channels")
	})
	// https://github.com/gin-gonic/gin#parameters-in-path
	// NOTE: /:xxx はじまりだとエラー。先頭に/channels足す
	g.POST("/channels/:cname/messages", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "%s %s\n", "POST", "chat/channels"+ctx.Param("cname")+"/messages")
	})
	g.GET("/channels/:cname/messages", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "%s %s\n", "GET", "chat/channels"+ctx.Param("cname")+"/messages")
	})

	h := router.Group("/contents")
	h.GET("", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "%s %s\n", "GET", "contents")
	})
	h.POST("", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "%s %s\n", "POST", "contents")
	})
	h.DELETE("", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "%s %s\n", "DELETE", "contents")
	})

	router.Run(":8080")
}
