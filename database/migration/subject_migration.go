package migration

import (
	"time"

	"github.com/edulink-api/database/migration/lib"
)

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type Subject struct {
	SubjectID       int64     `gorm:"primaryKey;autoIncrement"`
	GradeID         int64     `gorm:"not null"` // GradeID is the foreign key
	SubjectName     string    `gorm:"unique;not null"`
	SubjectDuration time.Time `gorm:"not null"`
	lib.BaseModel             /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (Subject) TableName() string {
	return lib.GenerateTableName(lib.Public, "subjects")
}
