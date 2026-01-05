package attendance

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

func (h *Handler) GetAllAttendance(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	search := c.Query("search", "")
	employeeID := c.Query("employee_id", "")

	offset := (page - 1) * limit

	var attendances []models.Attendance
	var total int64

	query := h.db.Model(&models.Attendance{}).Where("deleted = ?", 0)

	if search != "" {
		query = query.Where("attendance_comment LIKE ?", "%"+search+"%")
	}

	if employeeID != "" {
		query = query.Where("employee_id = ?", employeeID)
	}

	query.Count(&total).
		Offset(offset).
		Limit(limit).
		Find(&attendances)

	return c.JSON(fiber.Map{
		"data":  attendances,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func (h *Handler) GetAttendanceByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var attendance models.Attendance
	if err := h.db.First(&attendance, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Attendance not found"})
	}

	return c.JSON(attendance)
}

func (h *Handler) CreateAttendance(c *fiber.Ctx) error {
	var attendance models.Attendance
	if err := c.BodyParser(&attendance); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.db.Create(&attendance).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create attendance"})
	}

	return c.Status(201).JSON(attendance)
}

func (h *Handler) UpdateAttendance(c *fiber.Ctx) error {
	id := c.Params("id")

	var attendance models.Attendance
	if err := h.db.First(&attendance, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Attendance not found"})
	}

	if err := c.BodyParser(&attendance); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.db.Save(&attendance).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update attendance"})
	}

	return c.JSON(attendance)
}

func (h *Handler) DeleteAttendance(c *fiber.Ctx) error {
	id := c.Params("id")

	var attendance models.Attendance
	if err := h.db.First(&attendance, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Attendance not found"})
	}

	if err := h.db.Delete(&attendance).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete attendance"})
	}

	return c.JSON(fiber.Map{"message": "Attendance deleted successfully"})
}
