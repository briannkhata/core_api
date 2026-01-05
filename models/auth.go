package models

import "time"

// User represents tbl_users
type User struct {
	ID          int        `gorm:"primary_key" json:"id"`
	EmployeeID  *int       `json:"employee_id"`
	Username    string     `json:"username"`
	Password    string     `json:"password"`
	Deleted     int        `gorm:"default:0" json:"deleted"`
	DateAdded   *time.Time `json:"date_added"`
	AddedBy     *int       `json:"added_by"`
	LoginTrials int        `gorm:"default:5" json:"login_trials"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CreatedBy   *int       `json:"created_by"`
}

// Role represents tbl_roles
type Role struct {
	ID          int       `gorm:"primary_key" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	Deleted     int       `gorm:"default:0" json:"deleted"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   *int      `json:"created_by"`
	IsActive    int       `gorm:"default:1" json:"is_active"`
}

// UserRole represents tbl_user_roles
type UserRole struct {
	ID     int  `gorm:"primary_key" json:"id"`
	UserID *int `json:"user_id"`
	RoleID *int `json:"role_id"`
}

// Permission represents tbl_permissions
type Permission struct {
	ID          int       `gorm:"primary_key" json:"id"`
	Name        string    `json:"name"`
	MenuID      *int      `json:"menu_id"`
	Description string    `json:"description"`
	Deleted     int       `gorm:"default:0" json:"deleted"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   *int      `json:"created_by"`
}

// NavigationMenu represents tbl_navigation_menus
type NavigationMenu struct {
	ID          int       `gorm:"primary_key" json:"id"`
	Parent      string    `json:"parent"`
	Modu        string    `json:"modu"`
	Title       string    `gorm:"not null" json:"title"`
	URL         string    `json:"url"`
	Icon        string    `json:"icon"`
	Position    int       `gorm:"default:0" json:"position"`
	Deleted     int       `gorm:"default:0" json:"deleted"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   *int      `json:"created_by"`
	ModuleTitle string    `json:"module_title"`
	IsActive    int       `gorm:"default:1" json:"is_active"`
}

// ModuleRight represents tbl_module_rights
type ModuleRight struct {
	ID           int  `gorm:"primary_key" json:"id"`
	MenuID       *int `json:"menu_id"`
	RoleID       *int `json:"role_id"`
	PermissionID *int `json:"permission_id"`
}
