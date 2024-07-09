package migration

type AttendanceSummary struct {
	AttendanceSummaryID int64 `gorm:"primaryKey;autoIncrement"`
	StudentID int64
	TotalPresent int64
	TotalAlpha int64
	TotalPermission int64
	TotalSick int64
}
