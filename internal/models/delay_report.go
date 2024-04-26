package models

import "time"

type ReportStatus int8

const (
	ReportStatusUnknown ReportStatus = iota + 1
	ReportStatusPending
	ReportStatusResolved
)

type DelayReport struct {
	ID           int          `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID      int          `gorm:"null" json:"order_id"`
	AgentID      int          `gorm:"default:null" json:"agent_id"`
	ReportTime   time.Time    `gorm:"not null" json:"report_time"`
	ResolvedTime time.Time    `gorm:"null" json:"resolved_time"`
	Status       ReportStatus `gorm:"type:smallint;default:1" json:"status"`

	Order Order `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Agent Agent `gorm:"foreignKey:AgentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
