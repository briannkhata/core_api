package models

import (
	"time"
)

// Attendance represents tbl_attendances
type Attendance struct {
	BaseModel
	AttendanceDate    *time.Time `json:"attendance_date"`
	AttendanceDay     string     `json:"attendance_day"`
	IsWeekend         int        `gorm:"default:0" json:"is_weekend"`
	IsHoliday         int        `gorm:"default:0" json:"is_holiday"`
	AttendanceComment string     `json:"attendance_comment"`
	AttendanceCodeID  *int       `json:"attendance_code_id"`
	EmployeeID        *int       `json:"employee_id"`
	ShiftID           *int       `json:"shift_id"`
}
