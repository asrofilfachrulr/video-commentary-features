package routes

import (
	"comment/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(vr any) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Use(func(ctx *gin.Context) {
		ctx.Set("VideoRepository", vr)
	})

	r.POST("/video", handler.PostVideo)
	r.GET("/video", handler.GetVideo)

	return r
}
