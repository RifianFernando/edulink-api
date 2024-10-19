package migration

import "github.com/skripsi-be/database/migration/lib"

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type ClassName struct {
	ClassNameID   int64  `gorm:"primaryKey;autoIncrement"`
	GradeID       int64  `gorm:"not null"` // GradeID is the foreign key
	TeacherID     int64  `gorm:"not null"` // TeacherID is the foreign key
	Name          string `gorm:"type:char(1);not null"`
	lib.BaseModel        /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (ClassName) TableName() string {
	return lib.GenerateTableName(lib.Academic, "class_names")
}
