package models

import (
	"github.com/edulink-api/database/migration/lib"
)

type TeacherSubject struct {
	TeacherSubjectID int64 `gorm:"primaryKey" json:"id"`
	TeacherID        int64 `json:"teacher_id" binding:"required"`
	SubjectID        int64 `json:"subject_id" binding:"required"`
	lib.BaseModel
}

func (TeacherSubject) TableName() string {
	return lib.GenerateTableName(lib.Academic, "teacher_subjects")
}
