package migration

import "github.com/edulink-api/database/migration/lib"

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type ClassName struct {
	ClassNameID          int64                  `gorm:"primaryKey;autoIncrement"`
	GradeID              int64                  `gorm:"not null"` // Foreign key
	TeacherID            int64                  `gorm:"not null"` // Foreign key
	Name                 string                 `gorm:"type:char(1);not null"`
	Student              []Student              `gorm:"foreignKey:ClassNameID;references:ClassNameID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	TeachingClassSubject []TeachingClassSubject `gorm:"foreignKey:ClassNameID;references:ClassNameID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Score                []Score                `gorm:"foreignKey:ClassNameID;references:ClassNameID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	lib.BaseModel                               /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (ClassName) TableName() string {
	return lib.GenerateTableName(lib.Academic, "class_names")
}
