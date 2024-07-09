package migration

type TeacherSubject struct {
	TeacherID int64 `gorm:"primaryKey;autoIncrement"`
	SubjectID int64 `gorm:"not null"`
}
