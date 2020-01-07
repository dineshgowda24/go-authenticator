package models

import (
	"time"
)

//User struct declaration
type User struct {
	ID          uint
	CreatedAt   time.Time
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string `gorm:"type:varchar(100);unique_index"`
	Password    string
}

//Session struct declaration
type Session struct {
	ID            uint
	SessionXValue string    `gorm:"type:varchar(100);unique_index;not null"`
	UserID        uint      `sql:"not null"`
	ExpiresAt     time.Time `sql:"not null"`
}
