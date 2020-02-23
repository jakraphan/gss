package models

import "time"

type Article struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Title     string     `json:"title"`
	Excerpt   string     `json:"excerpt"`
	Body      string     `json:"body"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}
