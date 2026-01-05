package models

// StaffType represents tbl_staff_types
type StaffType struct {
	BaseModel
	Name         string `json:"name"`
	Days         *int   `json:"days"`
	DaysPerMonth *int   `json:"days_per_month"`
	HoursPerDay  *int   `json:"hours_per_day"`
}
