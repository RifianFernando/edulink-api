package models

import "github.com/edulink-api/database/migration/lib"

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type LearningSchedule struct {
	ScheduleID             int64 `json:"schedule_id" binding:"required" validate:"required"`
	TeachingClassSubjectID int64 `json:"teaching_class_id" binding:"required" validate:"required"`
}

type LearningScheduleModel struct {
	LearningSchedule
	TeachingClassSubject   TeachingClassSubjectModel `gorm:"foreignKey:TeachingClassSubjectID;references:TeachingClassSubjectID"`
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (LearningSchedule) TableName() string {
	return lib.GenerateTableName(lib.Administration, "learning_schedules")
}
