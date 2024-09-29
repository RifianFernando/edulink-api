package models

import (
	"github.com/skripsi-be/database/migration/lib"
)

type Class struct {
	ClassID    int64  `gorm:"primaryKey" json:"id"`
	TeacherID  int64  `json:"id_teacher" binding:"required"`
	ClassName  string `json:"name" binding:"required"`
	ClassGrade string `json:"grade" binding:"required"`
	lib.BaseModel
}

func (Class) TableName() string {
	return lib.GenerateTableName(lib.Public, "classes")
}
