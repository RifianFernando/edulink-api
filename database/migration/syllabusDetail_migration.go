package migration

type SyllabusDetail struct {
	SyllabusID int64 `gorm:"not null"`
	SyllabusDetailSession int32
	SyllabusDetailLearningObjective string `gorm:"not null"`
	SyllabusDetailLearningActivities string `gorm:"not null"`
}
