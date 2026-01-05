package routes

import (
	"yathuerp/services/gateway/config"
	"yathuerp/shared/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func Setup(app *fiber.App, cfg *config.Config) {
	// API versioning
	api := app.Group("/api/v1")

	// Employees microservice routes
	employees := api.Group("/employees")
	employees.All("/*", func(c *fiber.Ctx) error {
		url := cfg.Services["employees"] + c.Path() + "?" + c.QueryString()
		if err := proxy.Do(c, url); err != nil {
			return err
		}
		c.Response().Header.Del(fiber.HeaderXForwardedFor)
		return nil
	})

	// Payroll microservice routes (protected)
	payroll := api.Group("/payroll")
	payroll.Use(middleware.JWTAuth())
	payroll.All("/*", func(c *fiber.Ctx) error {
		url := cfg.Services["payroll"] + c.Path() + "?" + c.QueryString()
		if err := proxy.Do(c, url); err != nil {
			return err
		}
		c.Response().Header.Del(fiber.HeaderXForwardedFor)
		return nil
	})

	// Attendance microservice routes (protected)
	attendance := api.Group("/attendance")
	attendance.Use(middleware.JWTAuth())
	attendance.All("/*", func(c *fiber.Ctx) error {
		url := cfg.Services["attendance"] + c.Path() + "?" + c.QueryString()
		if err := proxy.Do(c, url); err != nil {
			return err
		}
		c.Response().Header.Del(fiber.HeaderXForwardedFor)
		return nil
	})

	// Leave Management microservice routes (protected)
	leave := api.Group("/leave")
	leave.Use(middleware.JWTAuth())
	leave.All("/*", func(c *fiber.Ctx) error {
		url := cfg.Services["leave"] + c.Path() + "?" + c.QueryString()
		if err := proxy.Do(c, url); err != nil {
			return err
		}
		c.Response().Header.Del(fiber.HeaderXForwardedFor)
		return nil
	})

	// Loan Management microservice routes (protected)
	loans := api.Group("/loans")
	loans.Use(middleware.JWTAuth())
	loans.All("/*", func(c *fiber.Ctx) error {
		url := cfg.Services["loans"] + c.Path() + "?" + c.QueryString()
		if err := proxy.Do(c, url); err != nil {
			return err
		}
		c.Response().Header.Del(fiber.HeaderXForwardedFor)
		return nil
	})

	// User Management microservice routes
	users := api.Group("/users")
	users.All("/*", func(c *fiber.Ctx) error {
		url := cfg.Services["users"] + c.Path() + "?" + c.QueryString()
		if err := proxy.Do(c, url); err != nil {
			return err
		}
		c.Response().Header.Del(fiber.HeaderXForwardedFor)
		return nil
	})

	// Performance Management microservice routes (protected)
	performance := api.Group("/performance")
	performance.Use(middleware.JWTAuth())
	performance.All("/*", func(c *fiber.Ctx) error {
		url := cfg.Services["performance"] + c.Path() + "?" + c.QueryString()
		if err := proxy.Do(c, url); err != nil {
			return err
		}
		c.Response().Header.Del(fiber.HeaderXForwardedFor)
		return nil
	})

	// Banking & Benefits microservice routes (protected)
	banking := api.Group("/banking")
	banking.Use(middleware.JWTAuth())
	banking.All("/*", func(c *fiber.Ctx) error {
		url := cfg.Services["banking"] + c.Path() + "?" + c.QueryString()
		if err := proxy.Do(c, url); err != nil {
			return err
		}
		c.Response().Header.Del(fiber.HeaderXForwardedFor)
		return nil
	})

	// Audit & Settings microservice routes (protected)
	audit := api.Group("/audit")
	audit.Use(middleware.JWTAuth())
	audit.All("/*", func(c *fiber.Ctx) error {
		url := cfg.Services["audit"] + c.Path() + "?" + c.QueryString()
		if err := proxy.Do(c, url); err != nil {
			return err
		}
		c.Response().Header.Del(fiber.HeaderXForwardedFor)
		return nil
	})
}
