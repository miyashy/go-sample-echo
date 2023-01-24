package task

type TaskStatus int

const (
	TODO TaskStatus = iota + 1
	IN_PROGRESS
	DONE
)
