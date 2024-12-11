package migration

import "github.com/edulink-api/database/migration/lib"

// type AssignmentType string

// const (
// 	Exam AssignmentType = "Exam"
// 	Task AssignmentType = "Task"
// )

// func (g *AssignmentType) Scan(value interface{}) error {
// 	*g = AssignmentType(value.([]byte))
// 	return nil
// }

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type Assignment struct {
	AssignmentID   int64   `gorm:"primaryKey;autoIncrement"`
	TypeAssignment string  `gorm:"not null;unique;type:VARCHAR(50)"`
	Score          []Score `gorm:"foreignKey:AssignmentID;references:AssignmentID"`
	lib.BaseModel          /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (Assignment) TableName() string {
	return lib.GenerateTableName(lib.Academic, "assignments")
}
