package handler

import (
	"comment/controller"
	"comment/model"
	"comment/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostComment(ctx *gin.Context) {
	if strID := ctx.Param("id"); strID == "" {
		return
	} else {
		if ID, err := strconv.ParseInt(strID, 10, 64); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid id",
			})
			return
		}
	}

	vr := ctx.MustGet("VideoRepository").(*repository.VideoRepository)

	var input model.Comment

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error binding JSON",
		})

		return
	}

	if err := controller.PostComment(&input, vr, ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error",
		})

	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "succed post comment",
	})
}

func GetCommentsByVideoId(ctx *gin.Context) {
	vr := ctx.MustGet("VideoRepository").(*repository.VideoRepository)

	data := controller.GetVideos(vr)

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}
