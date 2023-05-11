package activity

type ActivityRequest struct {
	Title string `gorm:"type:varchar" json:"title" validate:"required"`
	Email string `gorm:"type:varchar" json:"email"`
}

type UpdateRequest struct {
	Title string `gorm:"type:varchar" json:"title" validate:"required"`
}
