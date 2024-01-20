package models

type HealthStatus int

const (
	PendingStatus HealthStatus = iota
	ApprovedStatus
	FailedStatus
)

type Operation struct {
	Id           int
	AccountId    int
	ServiceId    int
	HealthStatus HealthStatus
	CreationDate string
	UpdateDate   string
	Cost         float64
}
