package employees

import (
	"strconv"
	"yathuerp/models"
	"yathuerp/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GetAllEmployees(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	search := c.Query("search", "")
	status := c.Query("status", "")

	offset := (page - 1) * limit

	var employees []models.Employee
	var total int64

	query := h.db.Model(&models.Employee{}).Preload("Department").Preload("Position")

	if search != "" {
		query = query.Where("first_name ILIKE ? OR last_name ILIKE ? OR employee_code ILIKE ? OR email ILIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	// Count total records
	query.Count(&total)

	// Get paginated results
	if err := query.Offset(offset).Limit(limit).Find(&employees).Error; err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to fetch employees")
	}

	meta := utils.PaginationMeta{
		Page:       page,
		PerPage:    limit,
		Total:      int(total),
		TotalPages: int((total + int64(limit) - 1) / int64(limit)),
	}

	return utils.SendPaginated(c, "Employees retrieved successfully", employees, meta)
}

func (h *Handler) GetEmployeeByID(c *fiber.Ctx) error {
	id := c.Params("id")

	employeeUUID, err := uuid.Parse(id)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid employee ID")
	}

	var employee models.Employee
	if err := h.db.Preload("Department").Preload("Position").Preload("Manager").
		First(&employee, "id = ?", employeeUUID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.SendError(c, fiber.StatusNotFound, "Employee not found")
		}
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to fetch employee")
	}

	return utils.SendSuccess(c, "Employee retrieved successfully", employee)
}

func (h *Handler) CreateEmployee(c *fiber.Ctx) error {
	var req models.Employee
	if err := c.BodyParser(&req); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	// Validate required fields
	if req.FirstName == "" || req.LastName == "" || req.Email == "" || req.EmployeeCode == "" {
		return utils.SendError(c, fiber.StatusBadRequest, "First name, last name, email, and employee code are required")
	}

	// Check if employee code already exists
	var existingEmployee models.Employee
	if err := h.db.Where("employee_code = ?", req.EmployeeCode).First(&existingEmployee).Error; err == nil {
		return utils.SendError(c, fiber.StatusConflict, "Employee code already exists")
	}

	// Check if email already exists
	if err := h.db.Where("email = ?", req.Email).First(&existingEmployee).Error; err == nil {
		return utils.SendError(c, fiber.StatusConflict, "Email already exists")
	}

	// Create employee
	if err := h.db.Create(&req).Error; err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to create employee")
	}

	// Fetch created employee with relationships
	h.db.Preload("Department").Preload("Position").First(&req, "id = ?", req.ID)

	return utils.SendSuccess(c, "Employee created successfully", req)
}

func (h *Handler) UpdateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	employeeUUID, err := uuid.Parse(id)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid employee ID")
	}

	var employee models.Employee
	if err := h.db.First(&employee, "id = ?", employeeUUID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.SendError(c, fiber.StatusNotFound, "Employee not found")
		}
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to fetch employee")
	}

	var req models.Employee
	if err := c.BodyParser(&req); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	// Update fields
	employee.FirstName = req.FirstName
	employee.LastName = req.LastName
	employee.Email = req.Email
	employee.Phone = req.Phone
	employee.DateOfBirth = req.DateOfBirth
	employee.Gender = req.Gender
	employee.Address = req.Address
	employee.City = req.City
	employee.State = req.State
	employee.Country = req.Country
	employee.PostalCode = req.PostalCode
	employee.DepartmentID = req.DepartmentID
	employee.PositionID = req.PositionID
	employee.ManagerID = req.ManagerID
	employee.Salary = req.Salary
	employee.Status = req.Status

	if err := h.db.Save(&employee).Error; err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to update employee")
	}

	// Fetch updated employee with relationships
	h.db.Preload("Department").Preload("Position").Preload("Manager").First(&employee, "id = ?", employee.ID)

	return utils.SendSuccess(c, "Employee updated successfully", employee)
}

func (h *Handler) DeleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	employeeUUID, err := uuid.Parse(id)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid employee ID")
	}

	var employee models.Employee
	if err := h.db.First(&employee, "id = ?", employeeUUID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.SendError(c, fiber.StatusNotFound, "Employee not found")
		}
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to fetch employee")
	}

	// Check if employee has subordinates
	var subordinateCount int64
	h.db.Model(&models.Employee{}).Where("manager_id = ?", employeeUUID).Count(&subordinateCount)
	if subordinateCount > 0 {
		return utils.SendError(c, fiber.StatusBadRequest, "Cannot delete employee with subordinates")
	}

	if err := h.db.Delete(&employee).Error; err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to delete employee")
	}

	return utils.SendSuccess(c, "Employee deleted successfully", nil)
}
