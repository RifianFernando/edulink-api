package migration

import (
	"github.com/edulink-api/database/migration/lib"
)

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type Subject struct {
	SubjectID          int64            `gorm:"primaryKey;autoIncrement"`
	SubjectName        string           `gorm:"not null;unique"`
	DurationPerSession int              `gorm:"not null"`
	DurationPerWeek    int              `gorm:"not null"`
	TeacherSubjects    []TeacherSubject `gorm:"foreignKey:SubjectID;references:SubjectID;constraint:OnUpdate:SET NULL,OnDelete:SET NULL"`
	lib.BaseModel                       /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (Subject) TableName() string {
	return lib.GenerateTableName(lib.Academic, "subjects")
}
