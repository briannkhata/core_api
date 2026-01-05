package models

import "time"

// Bank represents tbl_banks
type Bank struct {
	BaseModel
	Name     string `json:"name"`
	BankCode string `json:"bank_code"`
	Abbrev   string `json:"abbrev"`
}

// BankDetail represents tbl_bank_details
type BankDetail struct {
	BaseModel
	EmployeeID    *int   `json:"employee_id"`
	BankID        *int   `json:"bank_id"`
	AccountNumber string `json:"account_number"`
	AccountType   string `json:"account_type"`
	Branch        string `json:"branch"`
	City          string `json:"city"`
	DateOpened    string `json:"date_opened"`
	IsActive      int    `gorm:"default:1" json:"is_active"`
}

// SchemeType represents tbl_scheme_types
type SchemeType struct {
	BaseModel
	Name string `json:"name"`
}

// MembershipType represents tbl_membership_types
type MembershipType struct {
	BaseModel
	Name         string   `json:"name"`
	Charge       *float64 `json:"charge"`
	SchemeTypeID *int     `json:"scheme_type_id"`
}

// Dependant represents tbl_dependants
type Dependant struct {
	BaseModel
	Name             string   `json:"name"`
	EmployeeID       *int     `json:"employee_id"`
	CanDebit         int      `gorm:"default:0" json:"can_debit"`
	DebitAmount      *float64 `json:"debit_amount"`
	Dob              string   `json:"dob"`
	SchemeTypeID     *int     `json:"scheme_type_id"`
	MembershipNumber string   `json:"membership_number"`
	IsAdult          string   `json:"is_adult"`
	DateJoined       string   `json:"date_joined"`
	Gender           string   `json:"gender"`
	MembershipTypeID *int     `json:"membership_type_id"`
	RelationshipType string   `json:"relationship_type"`
	IDNo             string   `json:"id_no"`
	IDType           string   `json:"id_type"`
}

// Spouse represents tbl_spounses
type Spouse struct {
	BaseModel
	Name       string     `json:"name"`
	Dob        *time.Time `json:"dob"`
	EmployeeID *int       `json:"employee_id"`
	IDNo       string     `json:"id_no"`
	IDType     string     `json:"id_type"`
	Gender     string     `json:"gender"`
}
