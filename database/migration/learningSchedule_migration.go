package migration

type LearningSchedule struct {
	LearningScheduleID int64  `gorm:"primaryKey;autoIncrement"`
	ScheduleID         int64  `gorm:"not null"`
	SubjectID          int64  `gorm:"not null"`
	ClassID            int64  `gorm:"not null"`
	DayOfWeek          string `gorm:"not null"`
	Schedule           Schedule  `gorm:"foreignKey:ScheduleID;references:ScheduleID"`
}