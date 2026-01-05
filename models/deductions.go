package models

// DeductionType represents tbl_deduction_types
type DeductionType struct {
	BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
	Code        string `json:"code"`
	IsRecurring int    `gorm:"default:0" json:"is_recurring"`
	IsStatic    int    `gorm:"default:0" json:"is_static"`
}

// Deduction represents tbl_deductions
type Deduction struct {
	BaseModel
	EmployeeID      int     `gorm:"not null" json:"employee_id"`
	DeductionTypeID int     `gorm:"not null" json:"deduction_type_id"`
	Amount          float64 `gorm:"default:0.00" json:"amount"`
	PayrollID       *int    `json:"payroll_id"`
}
