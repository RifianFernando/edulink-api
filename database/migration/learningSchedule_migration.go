package migration

import "github.com/edulink-api/database/migration/lib"

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type LearningSchedule struct {
	ScheduleID             int64 `gorm:"not null;uniqueIndex:unique_learning_schedule"`
	TeachingClassSubjectID int64 `gorm:"not null;uniqueIndex:unique_learning_schedule"`
	lib.BaseModel                /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (LearningSchedule) TableName() string {
	return lib.GenerateTableName(lib.Administration, "learning_schedules")
}
