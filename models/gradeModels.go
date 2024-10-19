package models

import (
	"github.com/skripsi-be/database/migration/lib"
)

type Grade struct {
	GradeID int64 `gorm:"primaryKey" json:"id"`
	Grade   int   `json:"grade" binding:"required"`
	lib.BaseModel
}

func (Grade) TableName() string {
	return lib.GenerateTableName(lib.Academic, "grades")
}
