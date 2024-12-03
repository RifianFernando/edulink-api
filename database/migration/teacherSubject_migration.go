package migration

import "github.com/edulink-api/database/migration/lib"

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type TeacherSubject struct {
	TeacherSubjectID int64 `gorm:"primaryKey;autoIncrement"`
	TeacherID        int64 `gorm:"not null;uniqueIndex:unique_teacher_subject"`
	SubjectID        int64 `gorm:"not null;uniqueIndex:unique_teacher_subject"`

	lib.BaseModel /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (TeacherSubject) TableName() string {
	return lib.GenerateTableName(lib.Academic, "teacher_subjects")
}
