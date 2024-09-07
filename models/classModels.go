package models

import (
	"github.com/skripsi-be/database/migration/lib"

)

type Class struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	IDTeacher   uint           `json:"id_teacher" binding:"required"`
	Name        string         `json:"name" binding:"required"`
	Grade       int            `json:"grade" binding:"required"`
	lib.BaseModel
}
