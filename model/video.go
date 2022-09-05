package model

type Video struct {
	Name     string `json:"name" binding:"required"`
	Desc     string `json:"desc"`
	Comments []Comment
}
