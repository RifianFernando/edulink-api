package migration

import "github.com/edulink-api/database/migration/lib"

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type TeachingClassSubject struct {
	TeachingClassSubjectID int64 `gorm:"primaryKey;autoIncrement"` /* this is the same as `gorm:"primary_key;auto_increment"` */
	TeacherSubjectID       int64 `gorm:"not null;uniqueIndex:unique_teaching_class_subject"`
	ClassNameID            int64 `gorm:"not null;uniqueIndex:unique_teaching_class_subject"`
	AcademicYearID         int64 `gorm:"not null"`
	lib.BaseModel                /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (TeachingClassSubject) TableName() string {
	return lib.GenerateTableName(lib.Academic, "teaching_class_subjects")
}
