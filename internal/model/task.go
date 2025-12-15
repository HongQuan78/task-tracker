package model

import (
	"time"
)

type Task struct {
	Id          int
	Title       string
	CreatedAt   time.Time
	IsDeleted   bool
	IsCompleted bool
}

func NewTask(title string, id int) Task {
	return Task{
		Id:          id,
		Title:       title,
		CreatedAt:   time.Now(),
		IsCompleted: false,
		IsDeleted:   false,
	}
}
