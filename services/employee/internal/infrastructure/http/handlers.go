package http

import (
	"strconv"

	"yathuerp/services/employee-service/internal/application"
	"yathuerp/shared/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Handler struct {
	createEmployeeUseCase *application.CreateEmployeeUseCase
	logger                utils.Logger
}

func NewHandler(
	createEmployeeUseCase *application.CreateEmployeeUseCase,
	logger utils.Logger,
) *Handler {
	return &Handler{
		createEmployeeUseCase: createEmployeeUseCase,
		logger:                logger,
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

type PaginatedEmployeesResponse struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Data    []EmployeeResponse   `json:"data"`
	Meta    utils.PaginationMeta `json:"meta"`
}

func (h *Handler) CreateEmployee(c *fiber.Ctx) error {
	var req CreateEmployeeRequest
	if err := c.BodyParser(&req); err != nil {
		h.logger.Error("Invalid request body", "error", err)
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	// Validate request
	if err := utils.ValidateStruct(&req); err != nil {
		h.logger.Error("Validation failed", "error", err)
		return utils.SendError(c, fiber.StatusBadRequest, err.Error())
	}

	// Convert to application request
	appReq := &application.CreateEmployeeRequest{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		Phone:        req.Phone,
		DateOfBirth:  req.DateOfBirth,
		Gender:       req.Gender,
		Address:      req.Address,
		City:         req.City,
		State:        req.State,
		Country:      req.Country,
		PostalCode:   req.PostalCode,
		DepartmentID: req.DepartmentID,
		PositionID:   req.PositionID,
		ManagerID:    req.ManagerID,
		Salary:       req.Salary,
	}

	// Execute use case
	response, err := h.createEmployeeUseCase.Execute(c.Context(), appReq)
	if err != nil {
		h.logger.Error("Failed to create employee", "error", err)
		return utils.SendError(c, fiber.StatusInternalServerError, err.Error())
	}

	h.logger.Info("Employee created successfully", "employee_id", response.EmployeeID)
	return utils.SendSuccess(c, "Employee created successfully", response)
}

func (h *Handler) GetEmployeeByID(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.logger.Error("Invalid employee ID", "error", err, "id", idStr)
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid employee ID")
	}

	// TODO: Implement GetEmployeeByID use case
	// For now, return mock response
	employee := &EmployeeResponse{
		ID:        id.String(),
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Status:    "active",
	}

	h.logger.Info("Retrieved employee", "employee_id", id)
	return utils.SendSuccess(c, "Employee retrieved successfully", employee)
}

func (h *Handler) GetAllEmployees(c *fiber.Ctx) error {
	// Parse query parameters
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	search := c.Query("search", "")
	departmentID := c.Query("department_id", "")
	managerID := c.Query("manager_id", "")
	status := c.Query("status", "")

	// Parse UUIDs
	var deptID *uuid.UUID
	if departmentID != "" {
		if id, err := uuid.Parse(departmentID); err == nil {
			deptID = &id
		}
	}

	var mgrID *uuid.UUID
	if managerID != "" {
		if id, err := uuid.Parse(managerID); err == nil {
			mgrID = &id
		}
	}

	// TODO: Implement GetAllEmployees use case
	// For now, return mock response
	employees := []EmployeeResponse{
		{
			ID:        uuid.New().String(),
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
			Status:    "active",
		},
		{
			ID:        uuid.New().String(),
			FirstName: "Jane",
			LastName:  "Smith",
			Email:     "jane.smith@example.com",
			Status:    "active",
		},
	}

	// Mock pagination
	total := 2
	totalPages := (total + limit - 1) / limit
	meta := utils.PaginationMeta{
		Page:       page,
		PerPage:    limit,
		Total:      total,
		TotalPages: totalPages,
	}

	response := PaginatedEmployeesResponse{
		Success: true,
		Message: "Employees retrieved successfully",
		Data:    employees,
		Meta:    meta,
	}

	h.logger.Info("Retrieved employees", "count", len(employees), "page", page)
	return c.JSON(response)
}

func (h *Handler) UpdateEmployee(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.logger.Error("Invalid employee ID", "error", err, "id", idStr)
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid employee ID")
	}

	var req CreateEmployeeRequest
	if err := c.BodyParser(&req); err != nil {
		h.logger.Error("Invalid request body", "error", err)
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	// Validate request
	if err := utils.ValidateStruct(&req); err != nil {
		h.logger.Error("Validation failed", "error", err)
		return utils.SendError(c, fiber.StatusBadRequest, err.Error())
	}

	// TODO: Implement UpdateEmployee use case
	// For now, return success response
	response := map[string]interface{}{
		"employee_id": id.String(),
		"message":     "Employee updated successfully",
	}

	h.logger.Info("Employee updated", "employee_id", id)
	return utils.SendSuccess(c, "Employee updated successfully", response)
}

func (h *Handler) DeleteEmployee(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.logger.Error("Invalid employee ID", "error", err, "id", idStr)
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid employee ID")
	}

	// TODO: Implement DeleteEmployee use case
	// For now, return success response
	h.logger.Info("Employee deleted", "employee_id", id)
	return utils.SendSuccess(c, "Employee deleted successfully", nil)
}
