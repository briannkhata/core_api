package models

import (
	"time"
)

// Shift represents tbl_shifts
type Shift struct {
	BaseModel
	Title     string     `json:"title"`
	StartTime *time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
	IsActive  int        `gorm:"default:0" json:"is_active"`
}
