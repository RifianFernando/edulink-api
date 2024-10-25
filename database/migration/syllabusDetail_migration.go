package migration

import "github.com/edulink-api/database/migration/lib"

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type SyllabusDetail struct {
	SyllabusID                       int64 `gorm:"not null"`
	SyllabusDetailSession            int32
	SyllabusDetailLearningObjective  string `gorm:"not null"`
	SyllabusDetailLearningActivities string `gorm:"not null"`
	lib.BaseModel                           /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (SyllabusDetail) TableName() string {
	return lib.GenerateTableName(lib.Academic, "syllabus_details")
}
