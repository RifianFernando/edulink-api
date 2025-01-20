package models

import (
	"github.com/edulink-api/connections"
	"github.com/edulink-api/database/migration/lib"
)

type Assignment struct {
	AssignmentID   int64  `gorm:"primaryKey"`
	TypeAssignment string `json:"assignment_type" binding:"required" validate:"min=1,max=50"`
	// Score          []Score `gorm:"foreignKey:AssignmentID;references:AssignmentID"`
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

func (assignment *Assignment) GetAssignmentByType() (*Assignment, error) {
	result := connections.DB.First(&assignment, "type_assignment = ?", assignment.TypeAssignment).Scan(&assignment)
	if result.Error != nil {
		return assignment, result.Error
	}

	return assignment, nil
}

func (assignment *Assignment) GetAllAssignmentType() ([]Assignment, error) {
	var assignments []Assignment
	result := connections.DB.Find(&assignments)
	if result.Error != nil {
		return assignments, result.Error
	}

	return assignments, nil
}
