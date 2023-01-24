package value

import "github.com/google/uuid"

type TaskId string

func NewTaskId() TaskId {
	return TaskId(uuid.New().String())
}

func NewTaskIdWithHint(value string) (TaskId, error) {
	_, err := uuid.Parse(value)
	if err != nil {
		return TaskId(""), err
	}
	return TaskId(value), nil
}

func (taskId TaskId) Value() string {
	return string(taskId)
}
