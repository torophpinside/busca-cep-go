package model

import "time"

type Base struct {
	ID        uint       `gorm:"primary_key" sql:"id" json:"id"`
	CreatedAt *time.Time `sql:"created_at" json:"created"`
	UpdatedAt *time.Time `sql:"created_at" json:"updated"`
	DeletedAt *time.Time `sql:"created_at" json:"deleted"`
}
