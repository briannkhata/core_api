package models

import (
	"time"
)

// Setting represents tbl_settings
type Setting struct {
	ID                int       `gorm:"primary_key" json:"id"`
	Name              string    `json:"name"`
	PhysicalAddress   string    `json:"physical_address"`
	ContactAddress    string    `json:"contact_address"`
	Phone             string    `json:"phone"`
	AltPhone          string    `json:"alt_phone"`
	Email             string    `json:"email"`
	AltEmail          string    `json:"alt_email"`
	Country           string    `json:"country"`
	Logo              string    `json:"logo"`
	ContactPerson     string    `json:"contact_person"`
	OTCalculationMode string    `json:"ot_calculation_mode"`
	Deleted           int       `gorm:"default:0" json:"deleted"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	CreatedBy         *int      `json:"created_by"`
	DeductPayee       int       `gorm:"default:1" json:"deduct_payee"`
}

// Month represents tbl_months
type Month struct {
	Month     string    `gorm:"primary_key" json:"month"`
	Deleted   int       `gorm:"default:0" json:"deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy *int      `json:"created_by"`
}

// Year represents tbl_years
type Year struct {
	Year      string    `gorm:"primary_key" json:"year"`
	Deleted   int       `gorm:"default:0" json:"deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy *int      `json:"created_by"`
}

// OffenceType represents tbl_offence_types
type OffenceType struct {
	BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
	Punishment  string `json:"punishment"`
}
