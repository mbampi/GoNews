package models

import "time"

// Post struct
type Post struct {
	ID        uint64    `gorm:"primary_id;auto_increment;unique" json:"id"`
	Title     string    `gorm:"size:30;not null;" json:"title"`
	Content   string    `gorm:"type:text;" json:"content"`
	Author    Author    `gorm:"foreignkey:AuthorID" json:"author"`
	AuthorID  uint64    `gorm:"not null" json:"author_id"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}

// TODO: Validate function
