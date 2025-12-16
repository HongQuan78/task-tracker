package model

import (
	"time"
)

type Task struct {
	Id          int
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTask(description string, id int) Task {
	now := time.Now()
	return Task{
		Id:          id,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
