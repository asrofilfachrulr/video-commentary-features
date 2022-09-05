package repository

import "comment/model"

type VideoRepository struct {
	Obj []model.Video
}

func NewVideoRepository() *VideoRepository {
	return &VideoRepository{
		Obj: []model.Video{},
	}
}

func (vr *VideoRepository) PostVideo(video *model.Video) {
	vr.Obj = append(vr.Obj, *video)
}

func (vr *VideoRepository) GetVideos() []model.Video {
	return vr.Obj
}

func (vr *VideoRepository) GetVideoById(videoID int) *model.Video {
	return &vr.Obj[videoID]
}
