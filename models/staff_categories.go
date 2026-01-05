package models

// StaffCategory represents tbl_staff_categories
type StaffCategory struct {
	BaseModel
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Duration    *float64 `json:"duration"`
}
