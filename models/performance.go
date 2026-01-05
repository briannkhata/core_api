package models

import (
	"time"
)

// PerformanceCycle represents tbl_pf_cycles
type PerformanceCycle struct {
	ID                  int        `gorm:"primary_key" json:"id"`
	Name                string     `json:"name"`
	Description         string     `json:"description"`
	StartDate           time.Time  `json:"start_date"`
	EndDate             time.Time  `json:"end_date"`
	EvaluationStartDate *time.Time `json:"evaluation_start_date"`
	EvaluationEndDate   time.Time  `json:"evaluation_end_date"`
	Status              string     `gorm:"default:'draft'" json:"status"`
	CreatedBy           *int       `json:"created_by"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
}

// PerformanceAppraisal represents tbl_pf_appraisals
type PerformanceAppraisal struct {
	ID               int        `gorm:"primary_key" json:"id"`
	CycleID          int        `gorm:"not null" json:"cycle_id"`
	EmployeeID       int        `gorm:"not null" json:"employee_id"`
	ManagerID        int        `gorm:"not null" json:"manager_id"`
	Status           string     `gorm:"default:'draft'" json:"status"`
	OverallScore     *float64   `json:"overall_score"`
	OverallRating    string     `json:"overall_rating"`
	EmployeeComments string     `json:"employee_comments"`
	ManagerComments  string     `json:"manager_comments"`
	EmployeeSignedAt *time.Time `json:"employee_signed_at"`
	ManagerSignedAt  *time.Time `json:"manager_signed_at"`
	DiscussionDate   *time.Time `json:"discussion_date"`
	NextPeriodGoals  string     `json:"next_period_goals"`
	DevelopmentPlan  string     `json:"development_plan"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

// PerformanceGoal represents tbl_pf_goals
type PerformanceGoal struct {
	ID            int        `gorm:"primary_key" json:"id"`
	EmployeeID    int        `gorm:"not null" json:"employee_id"`
	CycleID       *int       `json:"cycle_id"`
	KpiID         *int       `json:"kpi_id"`
	Title         string     `gorm:"not null" json:"title"`
	Description   string     `json:"description"`
	TargetValue   *float64   `json:"target_value"`
	BaselineValue *float64   `json:"baseline_value"`
	CurrentValue  *float64   `json:"current_value"`
	Weight        float64    `gorm:"default:0.00" json:"weight"`
	Progress      float64    `gorm:"default:0.00" json:"progress"`
	DueDate       time.Time  `json:"due_date"`
	StartDate     *time.Time `json:"start_date"`
	Status        string     `gorm:"default:'active'" json:"status"`
	IsSmartGoal   int        `gorm:"default:0" json:"is_smart_goal"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

// PerformanceKPI represents tbl_pf_kpis
type PerformanceKPI struct {
	ID                       int       `gorm:"primary_key" json:"id"`
	DepartmentID             *int      `json:"department_id"`
	Title                    string    `gorm:"not null" json:"title"`
	Description              string    `json:"description"`
	Unit                     string    `gorm:"not null" json:"unit"`
	DataType                 string    `gorm:"default:'numeric'" json:"data_type"`
	TargetCalculationFormula string    `json:"target_calculation_formula"`
	IsActive                 int       `gorm:"default:1" json:"is_active"`
	CreatedBy                *int      `json:"created_by"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}

// PerformanceCompetency represents tbl_pf_competencies
type PerformanceCompetency struct {
	ID                   int       `gorm:"primary_key" json:"id"`
	ParentID             *int      `json:"parent_id"`
	Title                string    `gorm:"not null" json:"title"`
	Description          string    `json:"description"`
	CompetencyLevel      string    `gorm:"not null" json:"competency_level"`
	BehavioralIndicators string    `json:"behavioral_indicators"`
	Weight               float64   `gorm:"default:0.00" json:"weight"`
	IsActive             int       `gorm:"default:1" json:"is_active"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

// PerformancePIP represents tbl_pf_pips
type PerformancePIP struct {
	ID              int       `gorm:"primary_key" json:"id"`
	EmployeeID      int       `gorm:"not null" json:"employee_id"`
	AppraisalID     *int      `json:"appraisal_id"`
	Reason          string    `gorm:"not null" json:"reason"`
	Objectives      string    `gorm:"not null" json:"objectives"`
	SuccessCriteria string    `gorm:"not null" json:"success_criteria"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
	Status          string    `gorm:"default:'active'" json:"status"`
	Outcome         string    `json:"outcome"`
	CreatedBy       *int      `json:"created_by"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// PerformanceTraining represents tbl_pf_trainings
type PerformanceTraining struct {
	ID                   int        `gorm:"primary_key" json:"id"`
	Title                string     `gorm:"not null" json:"title"`
	Description          string     `json:"description"`
	TrainingType         string     `json:"training_type"`
	AssignedToEmployeeID *int       `json:"assigned_to_employee_id"`
	LinkedGoalID         *int       `json:"linked_goal_id"`
	LinkedCompetencyID   *int       `json:"linked_competency_id"`
	Status               string     `gorm:"default:'active'" json:"status"`
	StartDate            *time.Time `json:"start_date"`
	EndDate              *time.Time `json:"end_date"`
	Completed            int        `gorm:"default:0" json:"completed"`
	CompletionDate       *time.Time `json:"completion_date"`
	Cost                 *float64   `json:"cost"`
	Provider             string     `json:"provider"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
}
