package migration

import "github.com/skripsi-be/database/migration/lib"

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type Grade struct {
	GradeID       int64       `gorm:"primaryKey;autoIncrement"`
	Grade         int         `gorm:"not null"`
	Subjects      []Subject   `gorm:"foreignKey:GradeID;references:GradeID"`
	ClassNames    []ClassName `gorm:"foreignKey:GradeID;references:GradeID"`
	lib.BaseModel             /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (Grade) TableName() string {
	return lib.GenerateTableName(lib.Academic, "grades")
}
