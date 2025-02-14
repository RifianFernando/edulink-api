package models

import (
	"time"

	"github.com/edulink-api/database/migration/lib"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TeachingClassSubject struct {
	TeachingClassSubjectID int64 `gorm:"primaryKey"`
	TeacherSubjectID       int64 `gorm:"not null;uniqueIndex:unique_teaching_class_subject"`
	ClassNameID            int64 `gorm:"not null;uniqueIndex:unique_teaching_class_subject"`
	AcademicYearID         int64 `gorm:"not null"`
	lib.BaseModel
}

func (TeachingClassSubject) TableName() string {
	return lib.GenerateTableName(lib.Academic, "teaching_class_subjects")
}

// UpsertTeachingClassSubjects performs upsert on the teaching class subjects.
func UpsertTeachingClassSubjects(tx *gorm.DB, listTeacherSubjectIDIn []int64, records []TeachingClassSubject) error {
	// First, delete old records for the given teacher_subject_id list
	if err := tx.Where("teacher_subject_id IN ?", listTeacherSubjectIDIn).Delete(&TeachingClassSubject{}).Error; err != nil {
		return err
	}

	// Insert or update teaching class subjects
	return tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "teacher_subject_id"}, {Name: "class_name_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"deleted_at":       nil,
			"updated_at":       time.Now(),
			"academic_year_id": 1,
		}),
	}).Create(&records).Error
}
