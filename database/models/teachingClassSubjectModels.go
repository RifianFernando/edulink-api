package models

import (
	"github.com/edulink-api/database/migration/lib"
)

type TeachingClassSubject struct {
	TeachingClassSubjectID int64 `gorm:"primaryKey"`
	TeacherSubjectID       int64 `gorm:"not null;uniqueIndex:unique_teaching_class_subject"`
	ClassNameID            int64 `gorm:"not null;uniqueIndex:unique_teaching_class_subject"`
	lib.BaseModel
}

func (TeachingClassSubject) TableName() string {
	return lib.GenerateTableName(lib.Academic, "teaching_class_subjects")
}
