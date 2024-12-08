package models

import (
	"github.com/edulink-api/database/migration/lib"
)

type AcademicYear struct {
	AcademicYearID int64  `gorm:"primaryKey" json:"id"`
	AcademicYear   string `json:"academic_year" binding:"required"`
	lib.BaseModel
}

func (AcademicYear) TableName() string {
	return lib.GenerateTableName(lib.Academic, "academic_years")
}
