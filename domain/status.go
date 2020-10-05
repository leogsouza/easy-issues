package domain

// Status represents an issue status type
type Status string

// Type of status
var (
	StatusTodo       Status = "Todo"
	StatusInProgress Status = "In Progress"
	StatusDone       Status = "Done"
)
