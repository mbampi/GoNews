package models

import (
	"time"
	"errors"
)

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

// Validate validates the inputs
func (p *Post) Validate() error {
	if p.Title == "" {
		return errors.New("Title is required")
	}
	if p.Content == "" {
		return errors.New("Content is required")
	}
	if p.AuthorID < 0 {
		return errors.New("Author is required")
	}
	return nil
}
