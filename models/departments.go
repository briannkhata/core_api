package models

// Department represents tbl_departments
type Department struct {
	BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
}
