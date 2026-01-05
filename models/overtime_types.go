package models

// OvertimeType represents tbl_overtime_types
type OvertimeType struct {
	BaseModel
	Name        string  `json:"name"`
	Rate        float64 `json:"rate"`
	Code        string  `json:"code"`
	Description string  `json:"description"`
	IsCredit    int     `gorm:"default:1" json:"is_credit"`
}
