package models

import (
	"time"
)

// EmployeeGrade represents tbl_employee_grades
type EmployeeGrade struct {
	BaseModel
	EmployeeID       *int       `json:"employee_id"`
	StartDate        *time.Time `json:"start_date"`
	GradeID          *int       `json:"grade_id"`
	BasicSalary      *float64   `json:"basic_salary"`
	BranchID         *int       `json:"branch_id"`
	DepartmentID     *int       `json:"department_id"`
	JobID            *int       `json:"job_id"`
	ActionDate       *time.Time `json:"action_date"`
	IsCurrent        int        `gorm:"default:1" json:"is_current"`
	StaffTypeID      *int       `json:"staff_type_id"`
	Comment          string     `json:"comment"`
	ShiftType        string     `json:"shift_type"`
	EffectiveDate    *time.Time `json:"effective_date"`
	IncrementDetails string     `json:"increment_details"`
	StaffCategoryID  *int       `json:"staff_category_id"`
	EndDate          *time.Time `json:"end_date"`
}
