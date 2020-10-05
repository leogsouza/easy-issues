package domain

// Priority represents an issue importance
type Priority string

// Types of priorities
var (
	PriorityLow    Priority = "Low"
	PriorityMedium Priority = "Medium"
	PriorityHigh   Priority = "High"
)
