package migration

type Class struct {
	ClassID    int64  `gorm:"primaryKey;autoIncrement"`
	ClassName  string `gorm:"unique;not null"`
	ClassGrade string `gorm:"not null"`
	TeacherID  int64  `gorm:"foreignKey:TeacherID;references:TeacherID;constraint:OnUpdate:SET NULL,OnDelete:SET NULL"`
	Teacher    Teacher
}
