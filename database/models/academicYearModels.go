package models

import (
	"github.com/edulink-api/connections"
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

func (academicYear AcademicYear) CreateAcademicYear() error {
	result := connections.DB.Create(&academicYear)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (academicYear *AcademicYear) GetAcademicYearByModel() error {
	result := connections.DB.First(&academicYear,
		"academic_year = ?", academicYear.AcademicYear)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
