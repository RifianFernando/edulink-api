package migration

import "github.com/edulink-api/database/migration/lib"

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type Score struct {
	ScoreID        int64  `gorm:"primaryKey;autoIncrement"`
	StudentID      int64  `gorm:"not null"`
	AssignmentID   int64  `gorm:"not null"`
	TeacherID      int64  `gorm:"not null"`
	SubjectID      int64  `gorm:"not null"`
	AcademicYearID int64  `gorm:"not null"`
	Score          string `gorm:"not null"`
	lib.BaseModel         /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (Score) TableName() string {
	return lib.GenerateTableName(lib.Academic, "scores")
}
