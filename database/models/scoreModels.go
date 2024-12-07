package models

import (
	"github.com/edulink-api/connections"
	"github.com/edulink-api/database/migration/lib"
)

type Score struct {
	ScoreId        int64 `gorm:"primaryKey"`
	StudentID      int64 `json:"student_id" binding:"required"`
	AssignmentID   int64 `json:"assignment_id" binding:"required"`
	TeacherID      int64 `json:"teacher_id" binding:"required"`
	SubjectID      int64 `json:"subject_id" binding:"required"`
	AcademicYearID int64 `json:"academic_year_id" binding:"required"`
	Score          int   `json:"score" binding:"required"`
	lib.BaseModel
}

func (Score) TableName() string {
	return lib.GenerateTableName(lib.Administration, "Scores")
}

func GetAllScoringBySubjectClassID() (score []Score, err error) {
	result := connections.DB.Find(&score)
	if result.Error != nil {
		return []Score{}, result.Error
	} else if result.RowsAffected == 0 {
		return []Score{}, result.Error
	}

	return score, nil
}
