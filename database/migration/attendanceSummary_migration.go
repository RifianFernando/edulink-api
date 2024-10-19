package migration

import "github.com/skripsi-be/database/migration/lib"

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type AttendanceSummary struct {
	AttendanceSummaryID int64 `gorm:"primaryKey;autoIncrement"`
	AcademicYearID      int64 `gorm:"not null"`
	StudentID           int64
	TotalPresent        int
	TotalAlpha          int
	TotalPermission     int
	TotalSick           int
	lib.BaseModel       /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (AttendanceSummary) TableName() string {
	return lib.GenerateTableName(lib.Administration, "attendance_summaries")
}
