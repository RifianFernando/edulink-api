package migration

import (
	"time"

	"github.com/skripsi-be/database/migration/lib"
)

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
 */
type Subject struct {
	SubjectID       int64     `gorm:"primaryKey;autoIncrement"`
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
