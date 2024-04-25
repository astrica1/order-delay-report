package models

import "time"

type ReportStatus int8

const (
	ReportStatusUnknown ReportStatus = iota
	ReportStatusPending
	ReportStatusResolved
)

type DelayReport struct {
	ID           int          `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID      int          `gorm:"not null" json:"order_id"`
	AgentID      int          `gorm:"not null" json:"agent_id"`
	ReportTime   time.Time    `gorm:"not null" json:"report_time"`
	ResolvedTime time.Time    `gorm:"null" json:"resolved_time"`
	Status       ReportStatus `gorm:"null" json:"status"`

	Order Order `gorm:"foreignKey:OrderID"`
	Agent Agent `gorm:"foreignKey:AgentID"`
}
