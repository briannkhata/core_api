package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Employee struct {
	ID           uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	EmployeeCode string         `gorm:"uniqueIndex;not null" json:"employee_code"`
	FirstName    string         `gorm:"not null" json:"first_name"`
	LastName     string         `gorm:"not null" json:"last_name"`
	Email        string         `gorm:"uniqueIndex;not null" json:"email"`
	Phone        string         `json:"phone"`
	DateOfBirth  *time.Time     `json:"date_of_birth"`
	Gender       string         `json:"gender"`
	Address      string         `json:"address"`
	City         string         `json:"city"`
	State        string         `json:"state"`
	Country      string         `json:"country"`
	PostalCode   string         `json:"postal_code"`
	DepartmentID *uuid.UUID     `gorm:"type:uuid" json:"department_id"`
	PositionID   *uuid.UUID     `gorm:"type:uuid" json:"position_id"`
	ManagerID    *uuid.UUID     `gorm:"type:uuid" json:"manager_id"`
	HireDate     time.Time      `gorm:"not null" json:"hire_date"`
	Salary       float64        `json:"salary"`
	Status       string         `gorm:"default:'active'" json:"status"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Department   *Department `gorm:"foreignKey:DepartmentID" json:"department,omitempty"`
	Position     *Position   `gorm:"foreignKey:PositionID" json:"position,omitempty"`
	Manager      *Employee   `gorm:"foreignKey:ManagerID" json:"manager,omitempty"`
	Subordinates []Employee  `gorm:"foreignKey:ManagerID" json:"subordinates,omitempty"`
}

type Department struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Code        string         `gorm:"uniqueIndex;not null" json:"code"`
	Description string         `json:"description"`
	ParentID    *uuid.UUID     `gorm:"type:uuid" json:"parent_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Parent    *Department  `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Children  []Department `gorm:"foreignKey:ParentID" json:"children,omitempty"`
	Employees []Employee   `gorm:"foreignKey:DepartmentID" json:"employees,omitempty"`
}

type Position struct {
	ID           uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Title        string         `gorm:"not null" json:"title"`
	Code         string         `gorm:"uniqueIndex;not null" json:"code"`
	Description  string         `json:"description"`
	Level        int            `json:"level"`
	DepartmentID *uuid.UUID     `gorm:"type:uuid" json:"department_id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Department *Department `gorm:"foreignKey:DepartmentID" json:"department,omitempty"`
	Employees  []Employee  `gorm:"foreignKey:PositionID" json:"employees,omitempty"`
}
