package models

// Grade represents tbl_grades
type Grade struct {
	BaseModel
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	AnnualLeaveDays *int     `json:"annual_leave_days"`
	LeaveGrant      *float64 `json:"leave_grant"`
	EntrySalary     *float64 `json:"entry_salary"`
}
