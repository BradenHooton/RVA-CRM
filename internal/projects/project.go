package projects

import (
	"time"
	"rva_crm/internal/core"
)

type Project struct {
	core.BaseModel

	ProjectName string
	ProjectDescription string
	ProjectStatus string
	ProjectStartDate time.Time
	ProjectEndDate time.Time
	ProjectBudget float64
	ProjectProgress float64
	ProjectNotes string
	ProjectTasks []ProjectTask
}

type ProjectTask struct {
	core.BaseModel

	Project Project
	TaskName string
	TaskDescription string
	TaskStatus string
	TaskStartDate time.Time
	TaskEndDate time.Time
    Assignee string
    TaskType string
    TaskPriority string
}