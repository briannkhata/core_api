package payroll

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

func (h *Handler) GetAllPayrolls(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	search := c.Query("search", "")
	status := c.Query("status", "")

	offset := (page - 1) * limit

	var payrolls []models.Payroll
	var total int64

	query := h.db.Model(&models.Payroll{}).Where("deleted = ?", 0)

	if search != "" {
		query = query.Where("title LIKE ?", "%"+search+"%")
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Count(&total).
		Offset(offset).
		Limit(limit).
		Find(&payrolls)

	return c.JSON(fiber.Map{
		"data":  payrolls,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func (h *Handler) GetPayrollByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var payroll models.Payroll
	if err := h.db.First(&payroll, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Payroll not found"})
	}

	return c.JSON(payroll)
}

func (h *Handler) CreatePayroll(c *fiber.Ctx) error {
	var payroll models.Payroll
	if err := c.BodyParser(&payroll); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.db.Create(&payroll).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create payroll"})
	}

	return c.Status(201).JSON(payroll)
}

func (h *Handler) UpdatePayroll(c *fiber.Ctx) error {
	id := c.Params("id")

	var payroll models.Payroll
	if err := h.db.First(&payroll, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Payroll not found"})
	}

	if err := c.BodyParser(&payroll); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.db.Save(&payroll).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update payroll"})
	}

	return c.JSON(payroll)
}

func (h *Handler) DeletePayroll(c *fiber.Ctx) error {
	id := c.Params("id")

	var payroll models.Payroll
	if err := h.db.First(&payroll, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Payroll not found"})
	}

	if err := h.db.Delete(&payroll).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete payroll"})
	}

	return c.JSON(fiber.Map{"message": "Payroll deleted successfully"})
}

func (h *Handler) GetSalariesByPayroll(c *fiber.Ctx) error {
	payrollID := c.Params("payrollId")

	var salaries []models.Salary
	if err := h.db.Where("payroll_id = ? AND deleted = ?", payrollID, 0).Find(&salaries).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch salaries"})
	}

	return c.JSON(salaries)
}
