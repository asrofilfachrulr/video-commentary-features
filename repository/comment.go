package repository

import "comment/model"

func (vr *VideoRepository) PostComment(videoID int, c *model.Comment) {
	vr.Obj[videoID].Comments = append(vr.Obj[videoID].Comments, *c)
}

func (vr *VideoRepository) GetComments(videoID int) []model.Comment {
	return vr.GetVideoById(videoID).Comments
}
