package models

import (
	"time"
)

// Employee represents tbl_employees
type Employee struct {
	BaseModel
	FirstName        string     `json:"first_name"`
	MiddleName       string     `json:"middle_name"`
	LastName         string     `json:"last_name"`
	Gender           string     `json:"gender"`
	MaritalStatus    string     `json:"marital_status"`
	NextOfKin        string     `json:"next_of_kin"`
	PermanentAddress string     `json:"permanent_address"`
	ContactAddress   string     `json:"contact_address"`
	Email            string     `json:"email"`
	AltEmail         string     `json:"alt_email"`
	Dob              string     `json:"dob"`
	Phone            string     `json:"phone"`
	AltPhone         string     `json:"alt_phone"`
	DateAdded        *time.Time `json:"date_added"`
	OnPension        int        `gorm:"default:1" json:"on_pension"`
	Photo            string     `json:"photo"`
	NationalID       string     `json:"national_id"`
	Username         string     `json:"username"`
	Password         string     `json:"password"`
	Role             string     `gorm:"not null" json:"role"`
	AddedBy          *int       `json:"added_by"`
	IsAdmin          int        `gorm:"default:0" json:"is_admin"`
}
