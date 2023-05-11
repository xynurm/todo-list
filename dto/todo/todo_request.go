package todo

type TodoRequest struct {
	Title           string `gorm:"type:varchar" json:"title" validate:"required"`
	ActivityGroupID int    `gorm:"type:int" json:"activity_group_id"`
	IsActive        bool   `json:"is_active"`
	Priority        string `json:"priority"`
}

type UpdateRequest struct {
	Title string `gorm:"type:varchar" json:"title" validate:"required"`
}
