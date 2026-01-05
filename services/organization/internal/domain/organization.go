package domain

import (
	"time"

	"github.com/google/uuid"
)

// Department represents organizational department
type Department struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	Code        string     `json:"code" db:"code"`
	Description string     `json:"description" db:"description"`
	ParentID    *uuid.UUID `json:"parent_id" db:"parent_id"`
	ManagerID   *uuid.UUID `json:"manager_id" db:"manager_id"`
	Status      string     `json:"status" db:"status"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	CreatedBy   *uuid.UUID `json:"created_by" db:"created_by"`
}

// Branch represents organizational branch/office
type Branch struct {
	ID         uuid.UUID  `json:"id" db:"id"`
	Name       string     `json:"name" db:"name"`
	Code       string     `json:"code" db:"code"`
	Address    string     `json:"address" db:"address"`
	City       string     `json:"city" db:"city"`
	State      string     `json:"state" db:"state"`
	Country    string     `json:"country" db:"country"`
	PostalCode string     `json:"postal_code" db:"postal_code"`
	Phone      string     `json:"phone" db:"phone"`
	Email      string     `json:"email" db:"email"`
	ManagerID  *uuid.UUID `json:"manager_id" db:"manager_id"`
	Status     string     `json:"status" db:"status"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	CreatedBy  *uuid.UUID `json:"created_by" db:"created_by"`
}

// Repository interfaces
type DepartmentRepository interface {
	Create(department *Department) error
	GetByID(id uuid.UUID) (*Department, error)
	GetAll(filter *DepartmentFilter) ([]*Department, error)
	Update(department *Department) error
	Delete(id uuid.UUID) error
	GetByParentID(parentID uuid.UUID) ([]*Department, error)
	GetByManagerID(managerID uuid.UUID) ([]*Department, error)
}

type BranchRepository interface {
	Create(branch *Branch) error
	GetByID(id uuid.UUID) (*Branch, error)
	GetAll(filter *BranchFilter) ([]*Branch, error)
	Update(branch *Branch) error
	Delete(id uuid.UUID) error
	GetByManagerID(managerID uuid.UUID) ([]*Branch, error)
}

// Filters
type DepartmentFilter struct {
	ParentID  *uuid.UUID
	ManagerID *uuid.UUID
	Status    string
	Search    string
	Limit     int
	Offset    int
}

type BranchFilter struct {
	ManagerID *uuid.UUID
	Status    string
	Search    string
	Limit     int
	Offset    int
}

// Events
type DepartmentCreatedEvent struct {
	DepartmentID uuid.UUID `json:"department_id"`
	Timestamp    time.Time `json:"timestamp"`
}

type DepartmentUpdatedEvent struct {
	DepartmentID uuid.UUID   `json:"department_id"`
	OldValues    *Department `json:"old_values"`
	NewValues    *Department `json:"new_values"`
	Timestamp    time.Time   `json:"timestamp"`
}

type BranchCreatedEvent struct {
	BranchID  uuid.UUID `json:"branch_id"`
	Timestamp time.Time `json:"timestamp"`
}

type BranchUpdatedEvent struct {
	BranchID  uuid.UUID `json:"branch_id"`
	OldValues *Branch   `json:"old_values"`
	NewValues *Branch   `json:"new_values"`
	Timestamp time.Time `json:"timestamp"`
}
