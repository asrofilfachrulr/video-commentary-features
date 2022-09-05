package model

import "time"

type Comment struct {
	time.Time
	Content string
}
