package models

import (
	"time"
	"gorm.io/gorm"
)

type Class struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	IDTeacher   uint           `json:"id_teacher" binding:"required"`
	Name        string         `json:"name" binding:"required"`
	Grade       int            `json:"grade" binding:"required"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
