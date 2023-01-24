package task

import "go-sample-echo/domain/object/value"

type Task struct {
	Id          value.TaskId
	UserId      value.UserId
	Status      TaskStatus
	Description string
}

func NewTask(id value.TaskId, userId value.UserId, description string) Task {
	return Task{
		Id:          id,
		UserId:      userId,
		Status:      TODO,
		Description: description,
	}
}
