package migration

/* 
	* see the documentation here
	* https://gorm.io/docs/data_types.html
*/
type AttendanceSummary struct {
	AttendanceSummaryID int64 `gorm:"primaryKey;autoIncrement"`
	StudentID int64
	TotalPresent int64
	TotalAlpha int64
	TotalPermission int64
	TotalSick int64
	BaseModel /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
	* see the documentation here about conventions
	* https://gorm.io/docs/conventions.html
*/
func (AttendanceSummary) TableName() string {
	return GenerateTableName(Administration, "attendance_summaries")
}
