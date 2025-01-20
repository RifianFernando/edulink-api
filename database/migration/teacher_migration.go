package migration

import "github.com/edulink-api/database/migration/lib"

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type Teacher struct {
	TeacherID      int64            `gorm:"primaryKey;autoIncrement"`
	UserID         int64            `gorm:"not null"`
	TeachingHour   int32            `gorm:"not null"`
	ClassNames     []ClassName      `gorm:"foreignKey:TeacherID;references:TeacherID"`
	TeacherSubject []TeacherSubject `gorm:"foreignKey:TeacherID;references:TeacherID;constraint:OnUpdate:SET NULL,OnDelete:SET NULL"`
	lib.BaseModel                   /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (Teacher) TableName() string {
	return lib.GenerateTableName(lib.Academic, "teachers")
}
