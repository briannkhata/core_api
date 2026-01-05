package models

// Overtime represents tbl_overtimes
type Overtime struct {
	BaseModel
	HourlyRate     float64 `gorm:"default:0" json:"hourly_rate"`
	Hours          int     `gorm:"default:0" json:"hours"`
	Amount         float64 `gorm:"default:0" json:"amount"`
	EmployeeID     int     `gorm:"default:0" json:"employee_id"`
	PayrollID      *int    `json:"payroll_id"`
	OvertimeTypeID *int    `json:"overtime_type_id"`
	Rate           float64 `json:"rate"`
	DailyRate      float64 `json:"daily_rate"`
	Days           float64 `json:"days"`
}
