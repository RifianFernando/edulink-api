package migration

import "github.com/skripsi-be/database/migration/lib"

type AssignmentType string

const (
	Exam AssignmentType = "Exam"
	Task AssignmentType = "Task"
)

func (g *AssignmentType) Scan(value interface{}) error {
	*g = AssignmentType(value.([]byte))
	return nil
}

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
 */
type Assignment struct {
	AssignmentID   int64          `gorm:"primaryKey;autoIncrement"`
	TypeAssignment AssignmentType `gorm:"not null"`
	Grade          []Grade        `gorm:"foreignKey:AssignmentID;references:AssignmentID"`
	lib.BaseModel                 /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (Assignment) TableName() string {
	return lib.GenerateTableName(lib.Academic, "assignments")
}
