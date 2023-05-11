package models

import "time"

type Activity struct {
	ActivityID int       `json:"id" gorm:"primary_key:auto_increment"`
	Title      string    `json:"title" gorm:"type: varchar(255)"`
	Email      string    `json:"email" gorm:"type: varchar(255)"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
