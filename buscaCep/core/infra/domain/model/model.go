package model

import "time"

type Base struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `sql:"created_at" json:"created"`
	UpdatedAt time.Time  `sql:"created_at" json:"updated"`
	DeletedAt *time.Time `sql:"deleted_at" json:"deleted"`
}
