package domain

import (
	"time"

	"github.com/google/uuid"
)

// Employee represents the core employee entity
type Employee struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	EmployeeCode string     `json:"employee_code" db:"employee_code"`
	FirstName    string     `json:"first_name" db:"first_name"`
	LastName     string     `json:"last_name" db:"last_name"`
	Email        string     `json:"email" db:"email"`
	Phone        string     `json:"phone" db:"phone"`
	DateOfBirth  *time.Time `json:"date_of_birth" db:"date_of_birth"`
	Gender       string     `json:"gender" db:"gender"`
	Address      string     `json:"address" db:"address"`
	City         string     `json:"city" db:"city"`
	State        string     `json:"state" db:"state"`
	Country      string     `json:"country" db:"country"`
	PostalCode   string     `json:"postal_code" db:"postal_code"`
	DepartmentID *uuid.UUID `json:"department_id" db:"department_id"`
	PositionID   *uuid.UUID `json:"position_id" db:"position_id"`
	ManagerID    *uuid.UUID `json:"manager_id" db:"manager_id"`
	HireDate     time.Time  `json:"hire_date" db:"hire_date"`
	Salary       float64    `json:"salary" db:"salary"`
	Status       string     `json:"status" db:"status"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	CreatedBy    *uuid.UUID `json:"created_by" db:"created_by"`
}

// Department represents organizational department
type Department struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	Code        string     `json:"code" db:"code"`
	Description string     `json:"description" db:"description"`
	ParentID    *uuid.UUID `json:"parent_id" db:"parent_id"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	CreatedBy   *uuid.UUID `json:"created_by" db:"created_by"`
}

// Position represents job position
type Position struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	Title        string     `json:"title" db:"title"`
	Code         string     `json:"code" db:"code"`
	Description  string     `json:"description" db:"description"`
	Level        int        `json:"level" db:"level"`
	DepartmentID *uuid.UUID `json:"department_id" db:"department_id"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	CreatedBy    *uuid.UUID `json:"created_by" db:"created_by"`
}

// Repository defines the interface for employee data access
type Repository interface {
	Create(employee *Employee) error
	GetByID(id uuid.UUID) (*Employee, error)
	GetAll(filter *Filter) ([]*Employee, error)
	Update(employee *Employee) error
	Delete(id uuid.UUID) error
	GetByDepartmentID(deptID uuid.UUID) ([]*Employee, error)
	GetByManagerID(managerID uuid.UUID) ([]*Employee, error)
}

// Filter represents query filters for employees
type Filter struct {
	DepartmentID *uuid.UUID
	ManagerID    *uuid.UUID
	Status       string
	Search       string
	Limit        int
	Offset       int
}

// Events
type EmployeeCreatedEvent struct {
	EmployeeID uuid.UUID `json:"employee_id"`
	Timestamp  time.Time `json:"timestamp"`
}

type EmployeeUpdatedEvent struct {
	EmployeeID uuid.UUID `json:"employee_id"`
	OldValues  *Employee `json:"old_values"`
	NewValues  *Employee `json:"new_values"`
	Timestamp  time.Time `json:"timestamp"`
}

type EmployeeTerminatedEvent struct {
	EmployeeID      uuid.UUID `json:"employee_id"`
	TerminationDate time.Time `json:"termination_date"`
	Reason          string    `json:"reason"`
	Timestamp       time.Time `json:"timestamp"`
}
