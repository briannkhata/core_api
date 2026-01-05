package dto

import "time"

// Data Transfer Objects for Employee API

type CreateEmployeeRequest struct {
	FirstName    string  `json:"first_name" validate:"required"`
	LastName     string  `json:"last_name" validate:"required"`
	Email        string  `json:"email" validate:"required,email"`
	Phone        string  `json:"phone"`
	DateOfBirth  string  `json:"date_of_birth"`
	Gender       string  `json:"gender"`
	Address      string  `json:"address"`
	City         string  `json:"city"`
	State        string  `json:"state"`
	Country      string  `json:"country"`
	PostalCode   string  `json:"postal_code"`
	DepartmentID string  `json:"department_id"`
	PositionID   string  `json:"position_id"`
	ManagerID    string  `json:"manager_id"`
	Salary       float64 `json:"salary" validate:"min=0"`
}

type UpdateEmployeeRequest struct {
	FirstName    *string  `json:"first_name"`
	LastName     *string  `json:"last_name"`
	Email        *string  `json:"email"`
	Phone        *string  `json:"phone"`
	DateOfBirth  *string  `json:"date_of_birth"`
	Gender       *string  `json:"gender"`
	Address      *string  `json:"address"`
	City         *string  `json:"city"`
	State        *string  `json:"state"`
	Country      *string  `json:"country"`
	PostalCode   *string  `json:"postal_code"`
	DepartmentID *string  `json:"department_id"`
	PositionID   *string  `json:"position_id"`
	ManagerID    *string  `json:"manager_id"`
	Salary       *float64 `json:"salary"`
	Status       *string  `json:"status"`
}

type EmployeeResponse struct {
	ID           string  `json:"id"`
	EmployeeCode string  `json:"employee_code"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Email        string  `json:"email"`
	Phone        string  `json:"phone"`
	DateOfBirth  *string `json:"date_of_birth"`
	Gender       string  `json:"gender"`
	Address      string  `json:"address"`
	City         string  `json:"city"`
	State        string  `json:"state"`
	Country      string  `json:"country"`
	PostalCode   string  `json:"postal_code"`
	DepartmentID *string `json:"department_id"`
	PositionID   *string `json:"position_id"`
	ManagerID    *string `json:"manager_id"`
	HireDate     string  `json:"hire_date"`
	Salary       float64 `json:"salary"`
	Status       string  `json:"status"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type EmployeeListResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Data    []EmployeeResponse `json:"data"`
}

type PaginatedEmployeesResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Data    []EmployeeResponse `json:"data"`
	Meta    PaginationMeta     `json:"meta"`
}

type PaginationMeta struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

// Helper function to format date
func FormatDate(t *time.Time) *string {
	if t == nil {
		return nil
	}
	formatted := t.Format("2006-01-02")
	return &formatted
}
