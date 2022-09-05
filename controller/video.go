package controller

import (
	"comment/model"
	"comment/repository"
)

func PostVideo(video *model.Video, vr *repository.VideoRepository) error {
	vr.PostVideo(video)

	return nil
}

func GetVideos(vr *repository.VideoRepository) []model.Video {
	return vr.GetVideos()
}
