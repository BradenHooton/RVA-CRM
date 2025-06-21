package activity

import (
	"time"
	"rva_crm/internal/core"
)

type Activity struct {
	core.BaseModel

	ActivityType string
	ActivityDescription string
	ActivityDate time.Time
	ActivityStatus string
	ActivityPriority string
	ActivityCategory string
	ActivityTags []string
	ActivityMetadata map[string]interface{}
}

type ActivityType string
const (
	ActivityTypeCall ActivityType = "call"
	ActivityTypeEmail ActivityType = "email"
	ActivityTypeMeeting ActivityType = "meeting"
	ActivityTypeTask ActivityType = "task"
	ActivityTypeNote ActivityType = "note"
	ActivityTypeOther ActivityType = "other"
)

type ActivityStatus string
const (
	ActivityStatusPending ActivityStatus = "pending"
	ActivityStatusPostponed ActivityStatus = "postponed"
	ActivityStatusRescheduled ActivityStatus = "rescheduled"
	ActivityStatusNotStarted ActivityStatus = "not_started"
	ActivityStatusInProgress ActivityStatus = "in_progress"
	ActivityStatusOnHold ActivityStatus = "on_hold"
	ActivityStatusDeferred ActivityStatus = "deferred"
	ActivityStatusCancelled ActivityStatus = "cancelled"
	ActivityStatusCompleted ActivityStatus = "completed"
	ActivityStatusFailed ActivityStatus = "failed"
	ActivityStatusSkipped ActivityStatus = "skipped"
	ActivityStatusAwaitingApproval ActivityStatus = "awaiting_approval"
	ActivityStatusAwaitingPayment ActivityStatus = "awaiting_payment"
)