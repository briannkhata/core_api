package models

import (
	"time"
)

// Holiday represents tbl_holidays
type Holiday struct {
	BaseModel
	Title       string     `json:"title"`
	HolidayDate *time.Time `json:"holiday_date"`
	Comment     string     `json:"comment"`
}
