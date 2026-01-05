package domain

import (
	"time"

	"github.com/google/uuid"
)

// LeaveApplication represents employee leave request
type LeaveApplication struct {
	ID            uuid.UUID  `json:"id" db:"id"`
	EmployeeID    uuid.UUID  `json:"employee_id" db:"employee_id"`
	LeaveTypeID   uuid.UUID  `json:"leave_type_id" db:"leave_type_id"`
	StartDate     time.Time  `json:"start_date" db:"start_date"`
	EndDate       time.Time  `json:"end_date" db:"end_date"`
	DaysRequested float64    `json:"days_requested" db:"days_requested"`
	DaysApproved  float64    `json:"days_approved" db:"days_approved"`
	Reason        string     `json:"reason" db:"reason"`
	Status        string     `json:"status" db:"status"` // pending, approved, rejected, cancelled
	ApproverID    *uuid.UUID `json:"approver_id" db:"approver_id"`
	ApprovalDate  *time.Time `json:"approval_date" db:"approval_date"`
	ApprovalNotes string     `json:"approval_notes" db:"approval_notes"`
	Attachments   string     `json:"attachments" db:"attachments"` // JSON array of file URLs
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
}

// LeaveType represents types of leave available
type LeaveType struct {
	ID                 uuid.UUID `json:"id" db:"id"`
	Name               string    `json:"name" db:"name"`
	Code               string    `json:"code" db:"code"`
	Description        string    `json:"description" db:"description"`
	DaysPerYear        float64   `json:"days_per_year" db:"days_per_year"`
	IsPaid             bool      `json:"is_paid" db:"is_paid"`
	RequiresApproval   bool      `json:"requires_approval" db:"requires_approval"`
	MaxConsecutiveDays int       `json:"max_consecutive_days" db:"max_consecutive_days"`
	IsCarryForward     bool      `json:"is_carry_forward" db:"is_carry_forward"`
	CarryForwardLimit  float64   `json:"carry_forward_limit" db:"carry_forward_limit"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
}

// LeaveDay represents individual leave days for tracking
type LeaveDay struct {
	ID                 uuid.UUID `json:"id" db:"id"`
	LeaveApplicationID uuid.UUID `json:"leave_application_id" db:"leave_application_id"`
	Date               time.Time `json:"date" db:"date"`
	Type               string    `json:"type" db:"type"`     // full_day, half_day_morning, half_day_afternoon
	Status             string    `json:"status" db:"status"` // approved, rejected, pending
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
}

// LeaveBalance represents employee leave balance
type LeaveBalance struct {
	ID           uuid.UUID `json:"id" db:"id"`
	EmployeeID   uuid.UUID `json:"employee_id" db:"employee_id"`
	LeaveTypeID  uuid.UUID `json:"leave_type_id" db:"leave_type_id"`
	TotalDays    float64   `json:"total_days" db:"total_days"`
	UsedDays     float64   `json:"used_days" db:"used_days"`
	BalanceDays  float64   `json:"balance_days" db:"balance_days"`
	CarryForward float64   `json:"carry_forward" db:"carry_forward"`
	Year         int       `json:"year" db:"year"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// Repository interfaces
type LeaveApplicationRepository interface {
	Create(application *LeaveApplication) error
	GetByID(id uuid.UUID) (*LeaveApplication, error)
	GetByEmployeeID(employeeID uuid.UUID) ([]*LeaveApplication, error)
	GetPendingApplications() ([]*LeaveApplication, error)
	GetAll(filter *LeaveApplicationFilter) ([]*LeaveApplication, error)
	Update(application *LeaveApplication) error
	Delete(id uuid.UUID) error
	GetByDateRange(employeeID uuid.UUID, startDate, endDate time.Time) ([]*LeaveApplication, error)
}

type LeaveTypeRepository interface {
	Create(leaveType *LeaveType) error
	GetByID(id uuid.UUID) (*LeaveType, error)
	GetAll() ([]*LeaveType, error)
	Update(leaveType *LeaveType) error
	Delete(id uuid.UUID) error
	GetByCode(code string) (*LeaveType, error)
}

type LeaveDayRepository interface {
	Create(leaveDay *LeaveDay) error
	GetByID(id uuid.UUID) (*LeaveDay, error)
	GetByLeaveApplicationID(leaveApplicationID uuid.UUID) ([]*LeaveDay, error)
	Update(leaveDay *LeaveDay) error
	Delete(id uuid.UUID) error
	DeleteByLeaveApplicationID(leaveApplicationID uuid.UUID) error
}

type LeaveBalanceRepository interface {
	Create(balance *LeaveBalance) error
	GetByID(id uuid.UUID) (*LeaveBalance, error)
	GetByEmployeeAndLeaveType(employeeID, leaveTypeID uuid.UUID, year int) (*LeaveBalance, error)
	GetByEmployeeID(employeeID uuid.UUID, year int) ([]*LeaveBalance, error)
	Update(balance *LeaveBalance) error
	Delete(id uuid.UUID) error
}

// Filters
type LeaveApplicationFilter struct {
	EmployeeID  *uuid.UUID
	LeaveTypeID *uuid.UUID
	Status      string
	StartDate   *time.Time
	EndDate     *time.Time
	Limit       int
	Offset      int
}

// Events
type LeaveApplicationSubmittedEvent struct {
	ApplicationID uuid.UUID `json:"application_id"`
	EmployeeID    uuid.UUID `json:"employee_id"`
	Timestamp     time.Time `json:"timestamp"`
}

type LeaveApplicationApprovedEvent struct {
	ApplicationID uuid.UUID `json:"application_id"`
	EmployeeID    uuid.UUID `json:"employee_id"`
	ApproverID    uuid.UUID `json:"approver_id"`
	Timestamp     time.Time `json:"timestamp"`
}

type LeaveApplicationRejectedEvent struct {
	ApplicationID uuid.UUID `json:"application_id"`
	EmployeeID    uuid.UUID `json:"employee_id"`
	ApproverID    uuid.UUID `json:"approver_id"`
	Reason        string    `json:"reason"`
	Timestamp     time.Time `json:"timestamp"`
}
