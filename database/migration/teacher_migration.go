package migration

type Teacher struct {
	TeacherID    int64 `gorm:"primaryKey;autoIncrement"`
	UserID       int64 `gorm:"not null"`
	TeachingHour int32 `gorm:"not null"`
	Classes      []Class `gorm:"foreignKey:TeacherID;references:TeacherID"`
	Grades       []Grade `gorm:"foreignKey:TeacherID;references:TeacherID;constraint:OnUpdate:SET NULL,OnDelete:SET NULL"`
}
