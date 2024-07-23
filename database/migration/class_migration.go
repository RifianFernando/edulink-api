package migration

import "github.com/skripsi-be/database/migration/lib"

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
 */
type Class struct {
	ClassID       int64  `gorm:"primaryKey;autoIncrement"`
	ClassName     string `gorm:"unique;not null"`
	ClassGrade    string `gorm:"not null"`
	TeacherID     int64  `gorm:"foreignKey:TeacherID;references:TeacherID;constraint:OnUpdate:SET NULL,OnDelete:SET NULL"`
	Student       Student
	lib.BaseModel /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (Class) TableName() string {
	return lib.GenerateTableName(lib.Public, "classes")
}
