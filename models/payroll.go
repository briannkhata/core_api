package models

import (
	"time"
)

// Payroll represents tbl_payrolls
type Payroll struct {
	BaseModel
	Title  string `json:"title"`
	Month  string `json:"month"`
	Year   string `json:"year"`
	Status int    `gorm:"default:0" json:"status"`
}

// Salary represents tbl_salaries
type Salary struct {
	ID                  int        `gorm:"primary_key" json:"id"`
	NetSalary           float64    `gorm:"default:0" json:"net_salary"`
	GlossSalary         float64    `gorm:"default:0" json:"gloss_salary"`
	TotalPayee          float64    `json:"total_payee"`
	StaffContribution   float64    `json:"staff_contribution"`
	CompanyContribution float64    `json:"company_contribution"`
	TotalPension        float64    `json:"total_pension"`
	PayrollID           *int       `json:"payroll_id"`
	EmployeeID          *int       `json:"employee_id"`
	DateAdded           *time.Time `json:"date_added"`
	AddedBy             *int       `json:"added_by"`
	Deleted             int        `gorm:"default:0" json:"deleted"`
	BasicSalary         *int       `json:"basic_salary"`
	TotalOvertime       *int       `json:"total_overtime"`
	TotalDeductions     *float64   `json:"total_deductions"`
	TotalEarnings       *float64   `json:"total_earnings"`
	LeaveGrant          *float64   `json:"leave_grant"`
	HealthBill          *float64   `json:"health_bill"`
	AbsentCharge        *float64   `json:"absent_charge"`
	TotalLoans          *float64   `json:"total_loans"`
	OtherEarnings       *float64   `json:"other_earnings"`
	OtherDeductions     *float64   `json:"other_deductions"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	CreatedBy           *int       `json:"created_by"`
	IncludesPayee       string     `json:"includes_payee"`
}

// TaxBand represents tbl_tax_bands
type TaxBand struct {
	BaseModel
	Name      string   `json:"name"`
	Band1Top  *float64 `json:"band1_top"`
	Band2Top  *float64 `json:"band2_top"`
	Band3Top  *float64 `json:"band3_top"`
	Band4Top  *float64 `json:"band4_top"`
	Band1Rate *float64 `json:"band1_rate"`
	Band2Rate *float64 `json:"band2_rate"`
	Band3Rate *float64 `json:"band3_rate"`
	Band4Rate *float64 `json:"band4_rate"`
	IsActive  int      `gorm:"default:0" json:"is_active"`
}

// PensionParameter represents tbl_pension_parameters
type PensionParameter struct {
	BaseModel
	Name                string `json:"name"`
	StaffContribution   string `json:"staff_contribution"`
	CompanyContribution string `json:"company_contribution"`
	IsActive            int    `gorm:"default:1" json:"is_active"`
}
