package models

import "time"

type Todo struct {
	TodoID          int       `json:"id" gorm:"primary_key:auto_increment"`
	Title           string    `json:"title" gorm:"type: varchar(30)"`
	ActivityGroupID int       `json:"activity_group_id" gorm:"type: int(15)"`
	Priority        string    `json:"priority" gorm:"type: varchar(50)"`
	IsActive        bool      `json:"is_active" gorm:"default:false"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
