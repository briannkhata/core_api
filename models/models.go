package models

// This file contains all model definitions with proper GORM table names
// All models follow the tbl_modulename naming convention from the MySQL schema

// Table name constants for reference
const (
	TableAttendanceCodes   = "tbl_attendance_codes"
	TableAttendances       = "tbl_attendances"
	TableAuditLogs         = "tbl_audit_logs"
	TableBankDetails       = "tbl_bank_details"
	TableBanks             = "tbl_banks"
	TableBranches          = "tbl_branches"
	TableDeductionTypes    = "tbl_deduction_types"
	TableDeductions        = "tbl_deductions"
	TableDepartments       = "tbl_departments"
	TableDependants        = "tbl_dependants"
	TableEarningTypes      = "tbl_earning_types"
	TableEarnings          = "tbl_earnings"
	TableEmployeeGrades    = "tbl_employee_grades"
	TableEmployeeTrash     = "tbl_employee_trash"
	TableEmployees         = "tbl_employees"
	TableFinancialYears    = "tbl_financial_years"
	TableGrades            = "tbl_grades"
	TableHolidays          = "tbl_holidays"
	TableJobs              = "tbl_jobs"
	TableLeaveApplications = "tbl_leave_applications"
	TableLeaveDays         = "tbl_leave_days"
	TableLeaveTypes        = "tbl_leave_types"
	TableLoanApplications  = "tbl_loan_applications"
	TableLoanPayments      = "tbl_loan_payments"
	TableLoanTypes         = "tbl_loan_types"
	TableMembershipTypes   = "tbl_membership_types"
	TableModuleRights      = "tbl_module_rights"
	TableMonths            = "tbl_months"
	TableNavigationMenus   = "tbl_navigation_menus"
	TableOffenceTypes      = "tbl_offence_types"
	TableOvertimeTypes     = "tbl_overtime_types"
	TableOvertimes         = "tbl_overtimes"
	TablePayrolls          = "tbl_payrolls"
	TablePensionParameters = "tbl_pension_parameters"
	TablePermissions       = "tbl_permissions"
	TableRoles             = "tbl_roles"
	TableSalaries          = "tbl_salaries"
	TableSchemeTypes       = "tbl_scheme_types"
	TableSettings          = "tbl_settings"
	TableShifts            = "tbl_shifts"
	TableSpouses           = "tbl_spounses"
	TableStaffCategories   = "tbl_staff_categories"
	TableStaffTypes        = "tbl_staff_types"
	TableTaxBands          = "tbl_tax_bands"
	TableUserRoles         = "tbl_user_roles"
	TableUsers             = "tbl_users"
	TableYears             = "tbl_years"
	// Performance management tables (tbl_pf_*)
	TablePerformanceCycles       = "tbl_pf_cycles"
	TablePerformanceAppraisals   = "tbl_pf_appraisals"
	TablePerformanceGoals        = "tbl_pf_goals"
	TablePerformanceKPIs         = "tbl_pf_kpis"
	TablePerformanceCompetencies = "tbl_pf_competencies"
	TablePerformancePIPs         = "tbl_pf_pips"
	TablePerformanceTrainings    = "tbl_pf_trainings"
)

// TableName method for each model to ensure correct table names
func (AttendanceCode) TableName() string        { return TableAttendanceCodes }
func (Attendance) TableName() string            { return TableAttendances }
func (AuditLog) TableName() string              { return TableAuditLogs }
func (BankDetail) TableName() string            { return TableBankDetails }
func (Bank) TableName() string                  { return TableBanks }
func (Branch) TableName() string                { return TableBranches }
func (DeductionType) TableName() string         { return TableDeductionTypes }
func (Deduction) TableName() string             { return TableDeductions }
func (Department) TableName() string            { return TableDepartments }
func (Employee) TableName() string              { return TableEmployees }
func (EmployeeGrade) TableName() string         { return TableEmployeeGrades }
func (EmployeeTrash) TableName() string         { return TableEmployeeTrash }
func (Grade) TableName() string                 { return TableGrades }
func (StaffCategory) TableName() string         { return TableStaffCategories }
func (StaffType) TableName() string             { return TableStaffTypes }
func (Holiday) TableName() string               { return TableHolidays }
func (Job) TableName() string                   { return TableJobs }
func (LeaveApplication) TableName() string      { return TableLeaveApplications }
func (LeaveDay) TableName() string              { return TableLeaveDays }
func (LeaveType) TableName() string             { return TableLeaveTypes }
func (LoanApplication) TableName() string       { return TableLoanApplications }
func (LoanPayment) TableName() string           { return TableLoanPayments }
func (LoanType) TableName() string              { return TableLoanTypes }
func (MembershipType) TableName() string        { return TableMembershipTypes }
func (ModuleRight) TableName() string           { return TableModuleRights }
func (Month) TableName() string                 { return TableMonths }
func (NavigationMenu) TableName() string        { return TableNavigationMenus }
func (OffenceType) TableName() string           { return TableOffenceTypes }
func (OvertimeType) TableName() string          { return TableOvertimeTypes }
func (Overtime) TableName() string              { return TableOvertimes }
func (Payroll) TableName() string               { return TablePayrolls }
func (PensionParameter) TableName() string      { return TablePensionParameters }
func (Permission) TableName() string            { return TablePermissions }
func (Role) TableName() string                  { return TableRoles }
func (Salary) TableName() string                { return TableSalaries }
func (SchemeType) TableName() string            { return TableSchemeTypes }
func (Setting) TableName() string               { return TableSettings }
func (Shift) TableName() string                 { return TableShifts }
func (Spouse) TableName() string                { return TableSpouses }
func (TaxBand) TableName() string               { return TableTaxBands }
func (UserRole) TableName() string              { return TableUserRoles }
func (User) TableName() string                  { return TableUsers }
func (Year) TableName() string                  { return TableYears }
func (PerformanceCycle) TableName() string      { return TablePerformanceCycles }
func (PerformanceAppraisal) TableName() string  { return TablePerformanceAppraisals }
func (PerformanceGoal) TableName() string       { return TablePerformanceGoals }
func (PerformanceKPI) TableName() string        { return TablePerformanceKPIs }
func (PerformanceCompetency) TableName() string { return TablePerformanceCompetencies }
func (PerformancePIP) TableName() string        { return TablePerformancePIPs }
func (PerformanceTraining) TableName() string   { return TablePerformanceTrainings }
