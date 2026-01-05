package application

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"yathuerp/services/employee-service/internal/domain"
	"yathuerp/shared/logger"

	"github.com/google/uuid"
)

type CreateEmployeeUseCase struct {
	employeeRepo domain.Repository
	logger       logger.Logger
}

func NewCreateEmployeeUseCase(
	employeeRepo domain.Repository,
	logger logger.Logger,
) *CreateEmployeeUseCase {
	return &CreateEmployeeUseCase{
		employeeRepo: employeeRepo,
		logger:       logger,
	}
}

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

type CreateEmployeeResponse struct {
	EmployeeID string `json:"employee_id"`
	Message    string `json:"message"`
}

func (uc *CreateEmployeeUseCase) Execute(
	ctx context.Context,
	req *CreateEmployeeRequest,
) (*CreateEmployeeResponse, error) {
	// Validate business rules
	if err := uc.validateBusinessRules(req); err != nil {
		return nil, fmt.Errorf("business validation failed: %w", err)
	}

	// Convert to domain entity
	employee := uc.toDomainEntity(req)

	// Check if employee code already exists
	if existing, _ := uc.employeeRepo.GetByEmployeeCode(employee.EmployeeCode); existing != nil {
		return nil, fmt.Errorf("employee with code %s already exists", employee.EmployeeCode)
	}

	// Check if email already exists
	if existing, _ := uc.employeeRepo.GetByEmail(employee.Email); existing != nil {
		return nil, fmt.Errorf("employee with email %s already exists", employee.Email)
	}

	// Create employee
	if err := uc.employeeRepo.Create(employee); err != nil {
		uc.logger.Error("Failed to create employee", "error", err)
		return nil, fmt.Errorf("failed to create employee: %w", err)
	}

	// Publish event
	event := domain.EmployeeCreatedEvent{
		EmployeeID: employee.ID,
		Timestamp:  time.Now(),
	}

	// TODO: Publish to event bus
	// uc.eventBus.Publish("employee.created", event)

	uc.logger.Info("Employee created successfully", "employee_id", employee.ID)

	return &CreateEmployeeResponse{
		EmployeeID: employee.ID.String(),
		Message:    "Employee created successfully",
	}, nil
}

func (uc *CreateEmployeeUseCase) validateBusinessRules(req *CreateEmployeeRequest) error {
	// Validate salary is positive
	if req.Salary <= 0 {
		return fmt.Errorf("salary must be positive")
	}

	// Validate email format
	if !isValidEmail(req.Email) {
		return fmt.Errorf("invalid email format")
	}

	// Validate required fields
	if req.FirstName == "" || req.LastName == "" || req.Email == "" {
		return fmt.Errorf("first name, last name, and email are required")
	}

	return nil
}

func (uc *CreateEmployeeUseCase) toDomainEntity(req *CreateEmployeeRequest) *domain.Employee {
	employee := &domain.Employee{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Salary:    req.Salary,
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Parse optional fields
	if req.DepartmentID != "" {
		if deptID, err := uuid.Parse(req.DepartmentID); err == nil {
			employee.DepartmentID = &deptID
		}
	}

	if req.PositionID != "" {
		if posID, err := uuid.Parse(req.PositionID); err == nil {
			employee.PositionID = &posID
		}
	}

	if req.ManagerID != "" {
		if mgrID, err := uuid.Parse(req.ManagerID); err == nil {
			employee.ManagerID = &mgrID
		}
	}

	if req.DateOfBirth != "" {
		if dob, err := time.Parse("2006-01-02", req.DateOfBirth); err == nil {
			employee.DateOfBirth = &dob
		}
	}

	// Generate employee code
	employee.EmployeeCode = uc.generateEmployeeCode(req.FirstName, req.LastName)
	employee.ID = uuid.New()

	return employee
}

func (uc *CreateEmployeeUseCase) generateEmployeeCode(firstName, lastName string) string {
	// Generate employee code based on name and timestamp
	initials := strings.ToUpper(string(firstName[0]) + string(lastName[0]))
	timestamp := time.Now().Format("20060102")
	return fmt.Sprintf("%s-%s", initials, timestamp)
}

func isValidEmail(email string) bool {
	// Simple email validation
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(emailRegex, email)
	return matched
}
