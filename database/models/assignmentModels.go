package models

import (
	"github.com/edulink-api/connections"
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

func (assignment *Assignment) CreateAssignmentType() (*Assignment, error) {
	result := connections.DB.Create(&assignment)
	if result.Error != nil {
		return assignment, result.Error
	}

	return assignment, nil
}

func (assignment *Assignment) GetAssignment() (*Assignment, error) {
	result := connections.DB.First(&assignment, "type_assignment = ?", assignment.TypeAssignment).Scan(&assignment)
	if result.Error != nil {
		return assignment, result.Error
	}

	return assignment, nil
}
