package migration

import "github.com/edulink-api/database/migration/lib"

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
//  TODO: fix this struct
type HourSchedule struct {
	HourScheduleID int64      `gorm:"primaryKey;autoIncrement"`
	StartHour      int        `gorm:"unique;not null"`
	EndHour        int        `gorm:"unique;not null"`
	Schedule       []Schedule `gorm:"foreignKey:HourScheduleID;references:HourScheduleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	lib.BaseModel             /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (HourSchedule) TableName() string {
	return lib.GenerateTableName(lib.Academic, "hour_schedules")
}
