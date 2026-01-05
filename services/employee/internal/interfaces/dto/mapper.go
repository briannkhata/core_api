package dto

import (
	"time"
	"yathu_erp/services/employee-service/internal/domain"

	"github.com/google/uuid"
)

// EmployeeMapper handles mapping between domain entities and DTOs
type EmployeeMapper struct{}

func (m *EmployeeMapper) ToDomainEntity(req *CreateEmployeeRequest) *domain.Employee {
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
	employee.EmployeeCode = m.generateEmployeeCode(req.FirstName, req.LastName)
	employee.ID = uuid.New()

	return employee
}

func (m *EmployeeMapper) ToDTO(employee *domain.Employee) *EmployeeResponse {
	return &EmployeeResponse{
		ID:           employee.ID.String(),
		EmployeeCode: employee.EmployeeCode,
		FirstName:    employee.FirstName,
		LastName:     employee.LastName,
		Email:        employee.Email,
		Phone:        employee.Phone,
		DateOfBirth:  FormatDate(employee.DateOfBirth),
		Gender:       employee.Gender,
		Address:      employee.Address,
		City:         employee.City,
		State:        employee.State,
		Country:      employee.Country,
		PostalCode:   employee.PostalCode,
		DepartmentID: m.formatUUID(employee.DepartmentID),
		PositionID:   m.formatUUID(employee.PositionID),
		ManagerID:    m.formatUUID(employee.ManagerID),
		HireDate:     employee.HireDate.Format("2006-01-02"),
		Salary:       employee.Salary,
		Status:       employee.Status,
		CreatedAt:    employee.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:    employee.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}

func (m *EmployeeMapper) ToDTOList(employees []*domain.Employee) []EmployeeResponse {
	employeeResponses := make([]EmployeeResponse, len(employees))

	for i, emp := range employees {
		employeeResponses[i] = m.ToDTO(emp)
	}

	return employeeResponses
}

func (m *EmployeeMapper) formatUUID(uuid *uuid.UUID) *string {
	if uuid == nil {
		return nil
	}
	str := uuid.String()
	return &str
}

func (m *EmployeeMapper) generateEmployeeCode(firstName, lastName string) string {
	// Generate employee code based on name and timestamp
	initials := string(firstName[0]) + string(lastName[0])
	timestamp := time.Now().Format("20060102")
	return initials + "-" + timestamp
}
