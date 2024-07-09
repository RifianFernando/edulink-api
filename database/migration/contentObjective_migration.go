package migration

type ContentObjective struct {
	SyllabusID int64 `gorm:"not null"`
	ContentObjectiveContentTopic string `gorm:"not null"`
	ContentObjectiveLearningObjective string `gorm:"not null"`
	ContentObjectiveDuration int32 `gorm:"not null"`
}
