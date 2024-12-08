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

type ScoreModel struct {
	Score
	Student Student `gorm:"foreignKey:StudentID;references:StudentID"`
	Subject Subject `gorm:"foreignKey:SubjectID;references:SubjectID"`
}

func (Score) TableName() string {
	return lib.GenerateTableName(lib.Academic, "scores")
}

type ScoringBySubjectClassName struct {
	StudentID      int64  `json:"student_id"`
	StudentName    string `json:"StudentName"`
	AssignmentID   int64  `json:"assignment_id"`
	TypeAssignment string `json:"type_assignment"`
	SubjectName    string `json:"subject_name"`
	Score          int    `json:"score"`
}

func GetAllScoringBySubjectClassID(subjectID, classNameID, teacherID string) ([]ScoringBySubjectClassName, error) {
	query := `
		SELECT 
			st.student_id,
			st.student_name,
			sc.assignment_id,
			a.type_assignment,
			su.subject_name,
			sc.score
		FROM academic.scores sc
		JOIN academic.students st ON sc.student_id = st.student_id
		JOIN academic.subjects su ON sc.subject_id = su.subject_id
		JOIN academic.assignments a ON sc.assignment_id = a.assignment_id
		WHERE st.class_name_id = ? AND sc.subject_id = ? AND sc.teacher_id = ?
	`

	var results []ScoringBySubjectClassName
	result := connections.DB.Raw(query, classNameID, subjectID, teacherID).Scan(&results)
	if result.Error != nil {
		return nil, result.Error
	}

	return results, nil
}
