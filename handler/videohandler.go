package handler

import (
	"comment/controller"
	"comment/model"
	"comment/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostVideo(ctx *gin.Context) {
	vr := ctx.MustGet("VideoRepository").(*repository.VideoRepository)

	var input model.Video

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error binding JSON",
		})

		return
	}

	if err := controller.PostVideo(&input, vr); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error",
		})

	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "succeed post video",
	})
}

func GetVideo(ctx *gin.Context) {
	vr := ctx.MustGet("VideoRepository").(*repository.VideoRepository)

	data := controller.GetVideos(vr)

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}
