package migration

import (
	"github.com/edulink-api/database/migration/lib"
)

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type DaySchedule struct {
	DayScheduleID int64      `gorm:"primaryKey;autoIncrement"`
	DayName       string     `gorm:"not null;unique;type:varchar(10)"`
	Schedule      []Schedule `gorm:"foreignKey:DayScheduleID;references:DayScheduleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	lib.BaseModel            /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (DaySchedule) TableName() string {
	return lib.GenerateTableName(lib.Academic, "day_schedules")
}
