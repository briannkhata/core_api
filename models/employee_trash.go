package models

import (
	"time"
)

// EmployeeTrash represents tbl_employee_trash
type EmployeeTrash struct {
	BaseModel
	EmployeeID       int        `gorm:"not null" json:"employee_id"`
	Action           string     `gorm:"not null" json:"action"`
	ActionReason     string     `gorm:"not null" json:"action_reason"`
	ActionDate       time.Time  `gorm:"default:current_timestamp()" json:"action_date"`
	ActivatedDate    *time.Time `json:"activated_date"`
	ActivationReason string     `json:"activation_reason"`
}
