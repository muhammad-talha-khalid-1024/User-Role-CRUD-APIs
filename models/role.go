package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string         `json:"name" gorm:"size:100;not null"`
	Type      string         `json:"type" gorm:"size:100;not null"`
	Status    int16          `json:"status" gorm:"default:1"`
	CreatedAt time.Time      `json:"createdAt" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"type:timestamp;"`
}
