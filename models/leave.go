package models

import (
	"time"
)

// LeaveType represents tbl_leave_types
type LeaveType struct {
	BaseModel
	Name                  string `json:"name"`
	Description           string `json:"description"`
	MaximumDays           *int   `json:"maximum_days"`
	LeaveGrantEntitlement int    `gorm:"default:0" json:"leave_grant_entitlement"`
}

// LeaveApplication represents tbl_leave_applications
type LeaveApplication struct {
	BaseModel
	DaysApplied             int        `gorm:"default:0" json:"days_applied"`
	EmployeeID              int        `gorm:"not null" json:"employee_id"`
	LeaveGrantAmount        float64    `gorm:"default:0" json:"leave_grant_amount"`
	StartDate               *time.Time `json:"start_date"`
	EndDate                 *time.Time `json:"end_date"`
	IsActive                int        `gorm:"default:0" json:"is_active"`
	DateApplied             *time.Time `json:"date_applied"`
	Comment                 string     `json:"comment"`
	LeaveTypeID             string     `json:"leave_type_id"`
	ApplicationStatus       string     `json:"application_status"`
	ApplicationStatusReason string     `json:"application_status_reason"`
	ApplicationStatusBy     *int       `json:"application_status_by"`
	LodgedBy                *int       `json:"lodged_by"`
}

// LeaveDay represents tbl_leave_days
type LeaveDay struct {
	BaseModel
	FinancialYearID  *int     `json:"financial_year_id"`
	LeaveDaysBalance *float64 `json:"leave_days_balance"`
	LeaveDaysUsed    *float64 `json:"leave_days_used"`
	EmployeeID       *int     `json:"employee_id"`
}

// FinancialYear represents tbl_financial_years
type FinancialYear struct {
	BaseModel
	Name      string     `json:"name"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
	AddedBy   *int       `json:"added_by"`
	IsActive  int        `gorm:"default:0" json:"is_active"`
}
