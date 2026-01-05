package models

import (
	"time"
)

// AuditLog represents tbl_audit_logs
type AuditLog struct {
	ID        int       `gorm:"primary_key" json:"id"`
	UserID    *int      `json:"user_id"`
	TableName string    `gorm:"not null" json:"table_name"`
	RecordID  *int      `json:"record_id"`
	Action    string    `gorm:"not null" json:"action"`
	OldData   string    `json:"old_data"`
	NewData   string    `json:"new_data"`
	IPAddress string    `json:"ip_address"`
	UserAgent string    `json:"user_agent"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	Deleted   int       `gorm:"default:0" json:"deleted"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy *int      `json:"created_by"`
}
