package models

// AttendanceCode represents tbl_attendance_codes
type AttendanceCode struct {
	BaseModel
	Code           string `json:"code"`
	Description    string `json:"description"`
	IsDebit        int    `gorm:"default:0" json:"is_debit"`
	OvertimeTypeID *int   `json:"overtime_type_id"`
}
