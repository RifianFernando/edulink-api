package migration


type DomainAchievement struct {
	SyllabusID int64 `gorm:"not null"`
	DomainAchievementLearningAchievement string `gorm:"not null"`
	DomainAchievementDomain string `gorm:"not null"`
}
