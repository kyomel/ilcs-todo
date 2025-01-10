package constant

// TaskStatus represents the possible status values for a task
type TaskStatus string

const (
	// TaskStatusPending represents a task that is not yet completed
	TaskStatusPending TaskStatus = "pending"
	// TaskStatusCompleted represents a task that has been completed
	TaskStatusCompleted TaskStatus = "completed"
)

// IsValidTaskStatus checks if a given status string is valid
func IsValidTaskStatus(status string) bool {
	switch TaskStatus(status) {
	case TaskStatusPending, TaskStatusCompleted:
		return true
	default:
		return false
	}
}
