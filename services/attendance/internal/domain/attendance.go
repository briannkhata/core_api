package domain

import (
	"time"

	"github.com/google/uuid"
)

// Attendance represents employee attendance record
type Attendance struct {
	ID             uuid.UUID  `json:"id" db:"id"`
	EmployeeID     uuid.UUID  `json:"employee_id" db:"employee_id"`
	AttendanceDate time.Time  `json:"attendance_date" db:"attendance_date"`
	CheckIn        *time.Time `json:"check_in" db:"check_in"`
	CheckOut       *time.Time `json:"check_out" db:"check_out"`
	WorkHours      float64    `json:"work_hours" db:"work_hours"`
	OvertimeHours  float64    `json:"overtime_hours" db:"overtime_hours"`
	Status         string     `json:"status" db:"status"` // present, absent, late, half_day
	Notes          string     `json:"notes" db:"notes"`
	ApprovedBy     *uuid.UUID `json:"approved_by" db:"approved_by"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
}

// AttendanceCode represents attendance status codes
type AttendanceCode struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Code        string    `json:"code" db:"code"`
	Description string    `json:"description" db:"description"`
	Type        string    `json:"type" db:"type"` // present, absent, leave, holiday
	Color       string    `json:"color" db:"color"`
	IsPaid      bool      `json:"is_paid" db:"is_paid"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Shift represents work shift
type Shift struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Code         string    `json:"code" db:"code"`
	StartTime    string    `json:"start_time" db:"start_time"` // HH:MM format
	EndTime      string    `json:"end_time" db:"end_time"`     // HH:MM format
	BreakTime    string    `json:"break_time" db:"break_time"` // HH:MM format
	WorkHours    float64   `json:"work_hours" db:"work_hours"`
	Description  string    `json:"description" db:"description"`
	IsNightShift bool      `json:"is_night_shift" db:"is_night_shift"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// Holiday represents company holidays
type Holiday struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Date        time.Time `json:"date" db:"date"`
	Type        string    `json:"type" db:"type"` // national, company, religious
	Description string    `json:"description" db:"description"`
	IsRecurring bool      `json:"is_recurring" db:"is_recurring"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Repository interfaces
type AttendanceRepository interface {
	Create(attendance *Attendance) error
	GetByID(id uuid.UUID) (*Attendance, error)
	GetByEmployeeID(employeeID uuid.UUID, startDate, endDate time.Time) ([]*Attendance, error)
	GetAll(filter *AttendanceFilter) ([]*Attendance, error)
	Update(attendance *Attendance) error
	Delete(id uuid.UUID) error
	GetByDate(date time.Time) ([]*Attendance, error)
}

type AttendanceCodeRepository interface {
	Create(code *AttendanceCode) error
	GetByID(id uuid.UUID) (*AttendanceCode, error)
	GetAll() ([]*AttendanceCode, error)
	Update(code *AttendanceCode) error
	Delete(id uuid.UUID) error
	GetByCode(code string) (*AttendanceCode, error)
}

type ShiftRepository interface {
	Create(shift *Shift) error
	GetByID(id uuid.UUID) (*Shift, error)
	GetAll() ([]*Shift, error)
	Update(shift *Shift) error
	Delete(id uuid.UUID) error
	GetByCode(code string) (*Shift, error)
}

type HolidayRepository interface {
	Create(holiday *Holiday) error
	GetByID(id uuid.UUID) (*Holiday, error)
	GetByYear(year int) ([]*Holiday, error)
	GetByDateRange(startDate, endDate time.Time) ([]*Holiday, error)
	GetAll() ([]*Holiday, error)
	Update(holiday *Holiday) error
	Delete(id uuid.UUID) error
}

// Filters
type AttendanceFilter struct {
	EmployeeID *uuid.UUID
	StartDate  *time.Time
	EndDate    *time.Time
	Status     string
	Limit      int
	Offset     int
}

// Events
type AttendanceRecordedEvent struct {
	AttendanceID uuid.UUID `json:"attendance_id"`
	EmployeeID   uuid.UUID `json:"employee_id"`
	Date         time.Time `json:"date"`
	Timestamp    time.Time `json:"timestamp"`
}

type AttendanceUpdatedEvent struct {
	AttendanceID uuid.UUID   `json:"attendance_id"`
	OldValues    *Attendance `json:"old_values"`
	NewValues    *Attendance `json:"new_values"`
	Timestamp    time.Time   `json:"timestamp"`
}
