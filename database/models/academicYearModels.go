package models

import (
	"fmt"

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
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("academic year not found")
	}

	return nil
}

func (academicYear *AcademicYear) GetAcademicYearList() ([]AcademicYear, error) {
	var academicYearList []AcademicYear
	result := connections.DB.Find(&academicYearList)
	if result.Error != nil {
		return academicYearList, result.Error
	} else if result.RowsAffected == 0 {
		return academicYearList, fmt.Errorf("no academic year found")
	}

	return academicYearList, nil
}
