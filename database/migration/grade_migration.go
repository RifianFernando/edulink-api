package migration

type Grade struct {
	StudentID  int64 `gorm:"not null"`
	AssignmentID int64 `gorm:"not null"`
	TeacherID  int64 `gorm:"not null"`
	SubjectID  int64 `gorm:"not null"`
	Grade      string `gorm:"not null"`
}
