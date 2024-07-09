package migration

type Staff struct {
	StaffID int64 `gorm:"primaryKey;autoIncrement"`
	UserID int64 `gorm:"not null"`
	TeachingHour int32 `gorm:"not null"`
}
