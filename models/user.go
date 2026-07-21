package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              int            `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName       string         `json:"firstName" gorm:"size:100;not null"`
	LastName        string         `json:"lastName" gorm:"size:100;not null"`
	Email           string         `json:"email" gorm:"uniqueIndex;size:255;not null"`
	EmailVerifiedAt *time.Time     `json:"emailVerifiedAt" gorm:"type:timestamp;default:null"`
	Password        string         `json:"-" gorm:"size:255;not null"`
	Status          int16          `json:"status" gorm:"default:1"`
	RememberToken   string         `json:"rememberToken" gorm:"size:255;default:null"`
	CreatedAt       time.Time      `json:"createdAt" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time      `json:"updatedAt" gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt       gorm.DeletedAt `json:"deletedAt" gorm:"type:timestamp;"`
}
