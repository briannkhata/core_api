package models

// Branch represents tbl_branches
type Branch struct {
	BaseModel
	Name                  string   `json:"name"`
	Description           string   `json:"description"`
	BranchNormalDailyRate *float64 `json:"branch_normal_daily_rate"`
	BranchPublicDailyRate *float64 `json:"branch_public_daily_rate"`
}
