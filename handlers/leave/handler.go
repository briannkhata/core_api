package leave

import (
	"strconv"
	"yathuerp/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GetAllLeaveApplications(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	search := c.Query("search", "")
	status := c.Query("status", "")
	employeeID := c.Query("employee_id", "")

	offset := (page - 1) * limit

	var applications []models.LeaveApplication
	var total int64

	query := h.db.Model(&models.LeaveApplication{}).Where("deleted = ?", 0)

	if search != "" {
		query = query.Where("comment LIKE ?", "%"+search+"%")
	}

	if status != "" {
		query = query.Where("application_status = ?", status)
	}

	if employeeID != "" {
		query = query.Where("employee_id = ?", employeeID)
	}

	query.Count(&total).
		Offset(offset).
		Limit(limit).
		Find(&applications)

	return c.JSON(fiber.Map{
		"data":  applications,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func (h *Handler) GetLeaveApplicationByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var application models.LeaveApplication
	if err := h.db.First(&application, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Leave application not found"})
	}

	return c.JSON(application)
}

func (h *Handler) CreateLeaveApplication(c *fiber.Ctx) error {
	var application models.LeaveApplication
	if err := c.BodyParser(&application); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.db.Create(&application).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create leave application"})
	}

	return c.Status(201).JSON(application)
}

func (h *Handler) UpdateLeaveApplication(c *fiber.Ctx) error {
	id := c.Params("id")

	var application models.LeaveApplication
	if err := h.db.First(&application, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Leave application not found"})
	}

	if err := c.BodyParser(&application); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.db.Save(&application).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update leave application"})
	}

	return c.JSON(application)
}

func (h *Handler) ApproveLeaveApplication(c *fiber.Ctx) error {
	id := c.Params("id")

	var application models.LeaveApplication
	if err := h.db.First(&application, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Leave application not found"})
	}

	application.ApplicationStatus = "approved"

	if err := h.db.Save(&application).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to approve leave application"})
	}

	return c.JSON(fiber.Map{"message": "Leave application approved successfully"})
}

func (h *Handler) RejectLeaveApplication(c *fiber.Ctx) error {
	id := c.Params("id")

	var application models.LeaveApplication
	if err := h.db.First(&application, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Leave application not found"})
	}

	application.ApplicationStatus = "rejected"

	if err := h.db.Save(&application).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to reject leave application"})
	}

	return c.JSON(fiber.Map{"message": "Leave application rejected successfully"})
}

func (h *Handler) GetLeaveTypes(c *fiber.Ctx) error {
	var leaveTypes []models.LeaveType
	if err := h.db.Where("deleted = ?", 0).Find(&leaveTypes).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch leave types"})
	}

	return c.JSON(leaveTypes)
}
