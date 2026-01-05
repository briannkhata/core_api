package postgres

import (
	"context"
	"fmt"
	"time"

	"yathuerp/services/employee-service/internal/domain"
	"yathuerp/shared/logger"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db     *pgxpool.Pool
	logger logger.Logger
}

func NewRepository(db *pgxpool.Pool, logger logger.Logger) domain.Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (r *repository) Create(employee *domain.Employee) error {
	query := `
		INSERT INTO employees (
			id, employee_code, first_name, last_name, email, phone, 
			date_of_birth, gender, address, city, state, country, 
			postal_code, department_id, position_id, manager_id, 
			hire_date, salary, status, created_at, updated_at, created_by
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18
		)`

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.db.Exec(ctx, query,
		employee.ID,
		employee.EmployeeCode,
		employee.FirstName,
		employee.LastName,
		employee.Email,
		employee.Phone,
		employee.DateOfBirth,
		employee.Gender,
		employee.Address,
		employee.City,
		employee.State,
		employee.Country,
		employee.PostalCode,
		employee.DepartmentID,
		employee.PositionID,
		employee.ManagerID,
		employee.HireDate,
		employee.Salary,
		employee.Status,
		employee.CreatedAt,
		employee.UpdatedAt,
		employee.CreatedBy,
	)

	if err != nil {
		r.logger.Error("Failed to create employee", "error", err)
		return fmt.Errorf("failed to create employee: %w", err)
	}

	r.logger.Info("Employee created successfully", "employee_id", employee.ID)
	return nil
}

func (r *repository) GetByID(id uuid.UUID) (*domain.Employee, error) {
	query := `
		SELECT 
			id, employee_code, first_name, last_name, email, phone, 
			date_of_birth, gender, address, city, state, country, 
			postal_code, department_id, position_id, manager_id, 
			hire_date, salary, status, created_at, updated_at, created_by
		FROM employees 
		WHERE id = $1 AND deleted = false
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	employee := &domain.Employee{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&employee.ID,
		&employee.EmployeeCode,
		&employee.FirstName,
		&employee.LastName,
		&employee.Email,
		&employee.Phone,
		&employee.DateOfBirth,
		&employee.Gender,
		&employee.Address,
		&employee.City,
		&employee.State,
		&employee.Country,
		&employee.PostalCode,
		&employee.DepartmentID,
		&employee.PositionID,
		&employee.ManagerID,
		&employee.HireDate,
		&employee.Salary,
		&employee.Status,
		&employee.CreatedAt,
		&employee.UpdatedAt,
		&employee.CreatedBy,
	)

	if err != nil {
		if err.Error() == "no rows" {
			return nil, fmt.Errorf("employee not found")
		}
		r.logger.Error("Failed to get employee by ID", "error", err, "employee_id", id)
		return nil, fmt.Errorf("failed to get employee: %w", err)
	}

	return employee, nil
}

func (r *repository) GetAll(filter *domain.Filter) ([]*domain.Employee, error) {
	query := `
		SELECT 
			id, employee_code, first_name, last_name, email, phone, 
			date_of_birth, gender, address, city, state, country, 
			postal_code, department_id, position_id, manager_id, 
			hire_date, salary, status, created_at, updated_at, created_by
		FROM employees 
		WHERE deleted = false
	`

	args := []interface{}{}
	argIndex := 1

	// Add filters
	if filter.DepartmentID != nil {
		query += " AND department_id = $" + fmt.Sprint(argIndex) + "$"
		args = append(args, filter.DepartmentID)
		argIndex++
	}

	if filter.ManagerID != nil {
		query += " AND manager_id = $" + fmt.Sprint(argIndex) + "$"
		args = append(args, filter.ManagerID)
		argIndex++
	}

	if filter.Status != "" {
		query += " AND status = $" + fmt.Sprint(argIndex) + "$"
		args = append(args, filter.Status)
		argIndex++
	}

	if filter.Search != "" {
		query += " AND (first_name ILIKE $" + fmt.Sprint(argIndex) + "$ OR last_name ILIKE $" + fmt.Sprint(argIndex+1) + "$ OR employee_code ILIKE $" + fmt.Sprint(argIndex+2) + "$ OR email ILIKE $" + fmt.Sprint(argIndex+3) + "$)"
		args = append(args, "%"+filter.Search+"%", "%"+filter.Search+"%", "%"+filter.Search+"%", "%"+filter.Search+"%")
		argIndex += 4
	}

	// Add pagination
	query += " ORDER BY created_at DESC LIMIT $" + fmt.Sprint(argIndex) + "$ OFFSET $" + fmt.Sprint(argIndex+1) + "$"
	args = append(args, filter.Limit, filter.Offset)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		r.logger.Error("Failed to query employees", "error", err)
		return nil, fmt.Errorf("failed to query employees: %w", err)
	}
	defer rows.Close()

	var employees []*domain.Employee
	for rows.Next() {
		employee := &domain.Employee{}
		err := rows.Scan(
			&employee.ID,
			&employee.EmployeeCode,
			&employee.FirstName,
			&employee.LastName,
			&employee.Email,
			&employee.Phone,
			&employee.DateOfBirth,
			&employee.Gender,
			&employee.Address,
			&employee.City,
			&employee.State,
			&employee.Country,
			&employee.PostalCode,
			&employee.DepartmentID,
			&employee.PositionID,
			&employee.ManagerID,
			&employee.HireDate,
			&employee.Salary,
			&employee.Status,
			&employee.CreatedAt,
			&employee.UpdatedAt,
			&employee.CreatedBy,
		)

		if err != nil {
			r.logger.Error("Failed to scan employee row", "error", err)
			continue
		}

		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Error after scanning employee rows", "error", err)
		return nil, fmt.Errorf("error after scanning rows: %w", err)
	}

	r.logger.Info("Retrieved employees", "count", len(employees))
	return employees, nil
}

func (r *repository) Update(employee *domain.Employee) error {
	query := `
		UPDATE employees SET
			first_name = $2, last_name = $3, email = $4, phone = $5,
			date_of_birth = $6, gender = $7, address = $8, city = $9,
			state = $10, country = $11, postal_code = $12,
			department_id = $13, position_id = $14, manager_id = $15,
			salary = $16, status = $17, updated_at = $18
		WHERE id = $1 AND deleted = false
	`

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.db.Exec(ctx, query,
		employee.FirstName,
		employee.LastName,
		employee.Email,
		employee.Phone,
		employee.DateOfBirth,
		employee.Gender,
		employee.Address,
		employee.City,
		employee.State,
		employee.Country,
		employee.PostalCode,
		employee.DepartmentID,
		employee.PositionID,
		employee.ManagerID,
		employee.Salary,
		employee.Status,
		time.Now(),
		employee.ID,
	)

	if err != nil {
		r.logger.Error("Failed to update employee", "error", err, "employee_id", employee.ID)
		return fmt.Errorf("failed to update employee: %w", err)
	}

	r.logger.Info("Employee updated successfully", "employee_id", employee.ID)
	return nil
}

func (r *repository) Delete(id uuid.UUID) error {
	query := `UPDATE employees SET deleted = true, updated_at = $1 WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.db.Exec(ctx, query, time.Now(), id)

	if err != nil {
		r.logger.Error("Failed to delete employee", "error", err, "employee_id", id)
		return fmt.Errorf("failed to delete employee: %w", err)
	}

	r.logger.Info("Employee deleted successfully", "employee_id", id)
	return nil
}

func (r *repository) GetByDepartmentID(deptID uuid.UUID) ([]*domain.Employee, error) {
	query := `
		SELECT 
			id, employee_code, first_name, last_name, email, phone, 
			date_of_birth, gender, address, city, state, country, 
			postal_code, department_id, position_id, manager_id, 
			hire_date, salary, status, created_at, updated_at, created_by
		FROM employees 
		WHERE department_id = $1 AND deleted = false
		ORDER BY first_name, last_name
	`

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := r.db.Query(ctx, query, deptID)
	if err != nil {
		r.logger.Error("Failed to query employees by department", "error", err, "department_id", deptID)
		return nil, fmt.Errorf("failed to query employees by department: %w", err)
	}
	defer rows.Close()

	var employees []*domain.Employee
	for rows.Next() {
		employee := &domain.Employee{}
		err := rows.Scan(
			&employee.ID,
			&employee.EmployeeCode,
			&employee.FirstName,
			&employee.LastName,
			&employee.Email,
			&employee.Phone,
			&employee.DateOfBirth,
			&employee.Gender,
			&employee.Address,
			&employee.City,
			&employee.State,
			&employee.Country,
			&employee.PostalCode,
			&employee.DepartmentID,
			&employee.PositionID,
			&employee.ManagerID,
			&employee.HireDate,
			&employee.Salary,
			&employee.Status,
			&employee.CreatedAt,
			&employee.UpdatedAt,
			&employee.CreatedBy,
		)

		if err != nil {
			r.logger.Error("Failed to scan employee row", "error", err)
			continue
		}

		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Error after scanning employee rows", "error", err)
		return nil, fmt.Errorf("error after scanning rows: %w", err)
	}

	r.logger.Info("Retrieved employees by department", "department_id", deptID, "count", len(employees))
	return employees, nil
}

func (r *repository) GetByManagerID(managerID uuid.UUID) ([]*domain.Employee, error) {
	query := `
		SELECT 
			id, employee_code, first_name, last_name, email, phone, 
			date_of_birth, gender, address, city, state, country, 
			postal_code, department_id, position_id, manager_id, 
			hire_date, salary, status, created_at, updated_at, created_by
		FROM employees 
		WHERE manager_id = $1 AND deleted = false
		ORDER BY first_name, last_name
	`

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := r.db.Query(ctx, query, managerID)
	if err != nil {
		r.logger.Error("Failed to query employees by manager", "error", err, "manager_id", managerID)
		return nil, fmt.Errorf("failed to query employees by manager: %w", err)
	}
	defer rows.Close()

	var employees []*domain.Employee
	for rows.Next() {
		employee := &domain.Employee{}
		err := rows.Scan(
			&employee.ID,
			&employee.EmployeeCode,
			&employee.FirstName,
			&employee.LastName,
			&employee.Email,
			&employee.Phone,
			&employee.DateOfBirth,
			&employee.Gender,
			&employee.Address,
			&employee.City,
			&employee.State,
			&employee.Country,
			&employee.PostalCode,
			&employee.DepartmentID,
			&employee.PositionID,
			&employee.ManagerID,
			&employee.HireDate,
			&employee.Salary,
			&employee.Status,
			&employee.CreatedAt,
			&employee.UpdatedAt,
			&employee.CreatedBy,
		)

		if err != nil {
			r.logger.Error("Failed to scan employee row", "error", err)
			continue
		}

		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Error after scanning employee rows", "error", err)
		return nil, fmt.Errorf("error after scanning rows: %w", err)
	}

	r.logger.Info("Retrieved employees by manager", "manager_id", managerID, "count", len(employees))
	return employees, nil
}

// Helper methods for specific lookups
func (r *repository) GetByEmployeeCode(code string) (*domain.Employee, error) {
	query := `SELECT id FROM employees WHERE employee_code = $1 AND deleted = false`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var id uuid.UUID
	err := r.db.QueryRow(ctx, query, code).Scan(&id)
	if err != nil {
		if err.Error() == "no rows" {
			return nil, nil
		}
		return nil, err
	}

	return r.GetByID(id)
}

func (r *repository) GetByEmail(email string) (*domain.Employee, error) {
	query := `SELECT id FROM employees WHERE email = $1 AND deleted = false`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var id uuid.UUID
	err := r.db.QueryRow(ctx, query, email).Scan(&id)
	if err != nil {
		if err.Error() == "no rows" {
			return nil, nil
		}
		return nil, err
	}

	return r.GetByID(id)
}
