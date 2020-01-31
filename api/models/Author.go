package models

import (
	"gonews/api/security"
	"time"
)

// Author struct
type Author struct {
	ID        uint64    `gorm:"primary_id;auto_increment;unique" json:"id"`
	Name      string    `gorm:"size:20;not null;unique" json:"name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:60;not null" json:"password,omitempty"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	Posts     []Post    `json:"posts,omitempty"`
}

// BeforeSave hash the author password (automatically called by GORM)
func (a *Author) BeforeSave() error {
	hashedPassword, err := security.Hash(a.Password)
	if err != nil {
		return err
	}
	a.Password = string(hashedPassword)
	return nil
}
