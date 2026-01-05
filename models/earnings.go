package models

// EarningType represents tbl_earning_types
type EarningType struct {
	BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
	IsTaxable   int    `gorm:"default:1" json:"is_taxable"`
	Code        string `json:"code"`
	IsRecurring int    `gorm:"default:0" json:"is_recurring"`
	IsStatic    int    `gorm:"default:0" json:"is_static"`
}

// Earning represents tbl_earnings
type Earning struct {
	BaseModel
	EmployeeID    int     `gorm:"not null" json:"employee_id"`
	EarningTypeID int     `gorm:"not null" json:"earning_type_id"`
	Amount        float64 `gorm:"default:0.00" json:"amount"`
	PayrollID     *int    `json:"payroll_id"`
}
