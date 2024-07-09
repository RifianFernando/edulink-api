package migration

type Syllabus struct {
	SyllabusID int64 `gorm:"primaryKey;autoIncrement"`
	SubjectID  int64 `gorm:"not null"`
	ContentObjectives []ContentObjective `gorm:"foreignKey:SyllabusID;references:SyllabusID"`
	DomainAchievements []DomainAchievement `gorm:"foreignKey:SyllabusID;references:SyllabusID"`
	SyllabusDetails []SyllabusDetail `gorm:"foreignKey:SyllabusID;references:SyllabusID"`
}
