package domain

import (
	"time"

	"github.com/google/uuid"
)

// LoanApplication represents employee loan request
type LoanApplication struct {
	ID               uuid.UUID  `json:"id" db:"id"`
	EmployeeID       uuid.UUID  `json:"employee_id" db:"employee_id"`
	LoanTypeID       uuid.UUID  `json:"loan_type_id" db:"loan_type_id"`
	Amount           float64    `json:"amount" db:"amount"`
	InterestRate     float64    `json:"interest_rate" db:"interest_rate"`
	TermMonths       int        `json:"term_months" db:"term_months"`
	MonthlyPayment   float64    `json:"monthly_payment" db:"monthly_payment"`
	Purpose          string     `json:"purpose" db:"purpose"`
	Status           string     `json:"status" db:"status"` // pending, approved, rejected, disbursed, paid, defaulted
	ApproverID       *uuid.UUID `json:"approver_id" db:"approver_id"`
	ApprovalDate     *time.Time `json:"approval_date" db:"approval_date"`
	DisbursementDate *time.Time `json:"disbursement_date" db:"disbursement_date"`
	CompletionDate   *time.Time `json:"completion_date" db:"completion_date"`
	ApprovalNotes    string     `json:"approval_notes" db:"approval_notes"`
	Attachments      string     `json:"attachments" db:"attachments"` // JSON array of file URLs
	CreatedAt        time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at" db:"updated_at"`
}

// LoanType represents types of loans available
type LoanType struct {
	ID                  uuid.UUID `json:"id" db:"id"`
	Name                string    `json:"name" db:"name"`
	Code                string    `json:"code" db:"code"`
	Description         string    `json:"description" db:"description"`
	MinAmount           float64   `json:"min_amount" db:"min_amount"`
	MaxAmount           float64   `json:"max_amount" db:"max_amount"`
	DefaultInterestRate float64   `json:"default_interest_rate" db:"default_interest_rate"`
	MinTermMonths       int       `json:"min_term_months" db:"min_term_months"`
	MaxTermMonths       int       `json:"max_term_months" db:"max_term_months"`
	RequiresGuarantor   bool      `json:"requires_guarantor" db:"requires_guarantor"`
	MaxActiveLoans      int       `json:"max_active_loans" db:"max_active_loans"`
	EligibilityCriteria string    `json:"eligibility_criteria" db:"eligibility_criteria"`
	IsActive            bool      `json:"is_active" db:"is_active"`
	CreatedAt           time.Time `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time `json:"updated_at" db:"updated_at"`
}

// LoanPayment represents loan installment payments
type LoanPayment struct {
	ID                uuid.UUID  `json:"id" db:"id"`
	LoanApplicationID uuid.UUID  `json:"loan_application_id" db:"loan_application_id"`
	PaymentNumber     int        `json:"payment_number" db:"payment_number"`
	DueDate           time.Time  `json:"due_date" db:"due_date"`
	PaymentDate       *time.Time `json:"payment_date" db:"payment_date"`
	AmountDue         float64    `json:"amount_due" db:"amount_due"`
	AmountPaid        float64    `json:"amount_paid" db:"amount_paid"`
	InterestAmount    float64    `json:"interest_amount" db:"interest_amount"`
	PrincipalAmount   float64    `json:"principal_amount" db:"principal_amount"`
	BalanceAmount     float64    `json:"balance_amount" db:"balance_amount"`
	Status            string     `json:"status" db:"status"` // pending, paid, overdue, partial
	PaymentMethod     string     `json:"payment_method" db:"payment_method"`
	ReferenceNumber   string     `json:"reference_number" db:"reference_number"`
	Notes             string     `json:"notes" db:"notes"`
	CreatedAt         time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at" db:"updated_at"`
}

// LoanGuarantor represents loan guarantor information
type LoanGuarantor struct {
	ID                uuid.UUID  `json:"id" db:"id"`
	LoanApplicationID uuid.UUID  `json:"loan_application_id" db:"loan_application_id"`
	GuarantorName     string     `json:"guarantor_name" db:"guarantor_name"`
	GuarantorEmail    string     `json:"guarantor_email" db:"guarantor_email"`
	GuarantorPhone    string     `json:"guarantor_phone" db:"guarantor_phone"`
	GuarantorAddress  string     `json:"guarantor_address" db:"guarantor_address"`
	GuarantorID       *uuid.UUID `json:"guarantor_id" db:"guarantor_id"` // If guarantor is employee
	Relationship      string     `json:"relationship" db:"relationship"`
	IsApproved        bool       `json:"is_approved" db:"is_approved"`
	ApprovalDate      *time.Time `json:"approval_date" db:"approval_date"`
	ApprovalNotes     string     `json:"approval_notes" db:"approval_notes"`
	CreatedAt         time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at" db:"updated_at"`
}

// Repository interfaces
type LoanApplicationRepository interface {
	Create(application *LoanApplication) error
	GetByID(id uuid.UUID) (*LoanApplication, error)
	GetByEmployeeID(employeeID uuid.UUID) ([]*LoanApplication, error)
	GetPendingApplications() ([]*LoanApplication, error)
	GetAll(filter *LoanApplicationFilter) ([]*LoanApplication, error)
	Update(application *LoanApplication) error
	Delete(id uuid.UUID) error
	GetActiveLoans(employeeID uuid.UUID) ([]*LoanApplication, error)
}

type LoanTypeRepository interface {
	Create(loanType *LoanType) error
	GetByID(id uuid.UUID) (*LoanType, error)
	GetAll() ([]*LoanType, error)
	GetActive() ([]*LoanType, error)
	Update(loanType *LoanType) error
	Delete(id uuid.UUID) error
	GetByCode(code string) (*LoanType, error)
}

type LoanPaymentRepository interface {
	Create(payment *LoanPayment) error
	GetByID(id uuid.UUID) (*LoanPayment, error)
	GetByLoanApplicationID(loanApplicationID uuid.UUID) ([]*LoanPayment, error)
	Update(payment *LoanPayment) error
	Delete(id uuid.UUID) error
	GetOverduePayments() ([]*LoanPayment, error)
	GetUpcomingPayments(days int) ([]*LoanPayment, error)
}

type LoanGuarantorRepository interface {
	Create(guarantor *LoanGuarantor) error
	GetByID(id uuid.UUID) (*LoanGuarantor, error)
	GetByLoanApplicationID(loanApplicationID uuid.UUID) ([]*LoanGuarantor, error)
	Update(guarantor *LoanGuarantor) error
	Delete(id uuid.UUID) error
}

// Filters
type LoanApplicationFilter struct {
	EmployeeID *uuid.UUID
	LoanTypeID *uuid.UUID
	Status     string
	StartDate  *time.Time
	EndDate    *time.Time
	Limit      int
	Offset     int
}

// Events
type LoanApplicationSubmittedEvent struct {
	ApplicationID uuid.UUID `json:"application_id"`
	EmployeeID    uuid.UUID `json:"employee_id"`
	Amount        float64   `json:"amount"`
	Timestamp     time.Time `json:"timestamp"`
}

type LoanApplicationApprovedEvent struct {
	ApplicationID uuid.UUID `json:"application_id"`
	EmployeeID    uuid.UUID `json:"employee_id"`
	ApproverID    uuid.UUID `json:"approver_id"`
	Amount        float64   `json:"amount"`
	Timestamp     time.Time `json:"timestamp"`
}

type LoanDisbursedEvent struct {
	ApplicationID    uuid.UUID `json:"application_id"`
	EmployeeID       uuid.UUID `json:"employee_id"`
	Amount           float64   `json:"amount"`
	DisbursementDate time.Time `json:"disbursement_date"`
	Timestamp        time.Time `json:"timestamp"`
}

type LoanPaymentReceivedEvent struct {
	PaymentID         uuid.UUID `json:"payment_id"`
	LoanApplicationID uuid.UUID `json:"loan_application_id"`
	EmployeeID        uuid.UUID `json:"employee_id"`
	Amount            float64   `json:"amount"`
	Timestamp         time.Time `json:"timestamp"`
}
