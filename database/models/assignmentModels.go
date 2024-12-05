package models

import (
	"github.com/edulink-api/database/migration/lib"
)

type Assignment struct {
	AssignmentId   int64  `gorm:"primaryKey"`
	TypeAssignment string `json:"assignment_type" binding:"required" validate:"min=1,max=50"`
	// Score          []Score `gorm:"foreignKey:AssignmentId;references:AssignmentId"`
	lib.BaseModel
}

func (Assignment) TableName() string {
	return lib.GenerateTableName(lib.Academic, "assignments")
}
