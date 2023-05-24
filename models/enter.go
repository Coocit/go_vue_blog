package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"-"`
}
