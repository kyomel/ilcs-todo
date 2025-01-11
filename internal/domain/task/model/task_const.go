package model

type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusCompleted TaskStatus = "completed"
)

func IsValidTaskStatus(status string) bool {
	switch TaskStatus(status) {
	case TaskStatusPending, TaskStatusCompleted:
		return true
	default:
		return false
	}
}
