package models

// Job represents tbl_jobs
type Job struct {
	BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
}
