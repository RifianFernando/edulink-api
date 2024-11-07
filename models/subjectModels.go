package models

import (
	"github.com/edulink-api/database/migration/lib"
)

type Subject struct {
	SubjectID              int64  `gorm:"primaryKey"`
	GradeID                int64  `json:"id_grade" binding:"required"`
	SubjectName            string `json:"name" binding:"required"`
	SubjectDurationMinutes int    `json:"duration" binding:"required" validate:"gte=0"`
	lib.BaseModel
}

type SubjectModel struct {
	Subject
	Teacher Teacher `gorm:"foreignKey:TeacherID;references:TeacherID"`
	Grade   Grade   `gorm:"foreignKey:GradeID;references:GradeID"`
}

func (Subject) TableName() string {
	return lib.GenerateTableName(lib.Public, "subjects")
}
