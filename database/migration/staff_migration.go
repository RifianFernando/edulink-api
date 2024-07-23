package migration

/* 
	* see the documentation here
	* https://gorm.io/docs/data_types.html
*/
type Staff struct {
	StaffID int64 `gorm:"primaryKey;autoIncrement"`
	UserID int64 `gorm:"not null"`
	TeachingHour int32 `gorm:"not null"`
	BaseModel /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
	* see the documentation here about conventions
	* https://gorm.io/docs/conventions.html
*/
func (Staff) TableName() string {
	return GenerateTableName(Administration, "staffs")
}
