package models

import (
	"time"
)

// LoanType represents tbl_loan_types
type LoanType struct {
	BaseModel
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	InterestRate float64  `gorm:"default:0.00" json:"interest_rate"`
	MaxAmount    *float64 `json:"max_amount"`
	IsActive     int      `gorm:"default:1" json:"is_active"`
}

// LoanApplication represents tbl_loan_applications
type LoanApplication struct {
	BaseModel
	AmountApplied           float64    `gorm:"default:0" json:"amount_applied"`
	ApplicationDate         *time.Time `json:"application_date"`
	AmountPayable           float64    `gorm:"default:0" json:"amount_payable"`
	AmountReturned          float64    `gorm:"default:0" json:"amount_returned"`
	Balance                 *float64   `json:"balance"`
	PaymentPeriod           float64    `gorm:"default:0" json:"payment_period"`
	PaymentRate             *float64   `json:"payment_rate"`
	EmployeeID              *int       `json:"employee_id"`
	DeductMonth             string     `json:"deduct_month"`
	DeductYear              string     `gorm:"default:'0'" json:"deduct_year"`
	LoanTypeID              *int       `json:"loan_type_id"`
	IsActive                int        `gorm:"default:1" json:"is_active"`
	ApplicationStatus       int        `gorm:"default:0" json:"application_status"`
	ApplicationStatusBy     *int       `json:"application_status_by"`
	ApplicationStatusReason string     `json:"application_status_reason"`
	LoanPurpose             string     `json:"loan_purpose"`
	WitnessOne              string     `json:"witness_one"`
	WitnessOneAddress       string     `json:"witness_one_address"`
	WitnessOnePhone         string     `json:"witness_one_phone"`
	WitnessTwo              string     `json:"witness_two"`
	WitnessTwoAddress       string     `json:"witness_two_address"`
	WitnessTwoPhone         string     `json:"witness_two_phone"`
}

// LoanPayment represents tbl_loan_payments
type LoanPayment struct {
	ID         int        `gorm:"primary_key" json:"id"`
	PayrollID  *int       `json:"payroll_id"`
	Amount     *int       `json:"amount"`
	EmployeeID *int       `json:"employee_id"`
	LoanID     *int       `json:"loan_id"`
	DatePaid   *time.Time `json:"date_paid"`
	Deleted    int        `gorm:"default:0" json:"deleted"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	CreatedBy  *int       `json:"created_by"`
}
