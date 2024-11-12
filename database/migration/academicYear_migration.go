package migration

import "github.com/edulink-api/database/migration/lib"

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type AcademicYear struct {
	AcademicYearID      int64               `gorm:"primaryKey;autoIncrement"`
	AcademicYear        string              `gorm:"not null;type:CHAR(4)"`
	Reports             []Report            `gorm:"foreignKey:AcademicYearID;references:AcademicYearID"`
	Scores              []Score             `gorm:"foreignKey:AcademicYearID;references:AcademicYearID"`
	AttendanceSummaries []AttendanceSummary `gorm:"foreignKey:AcademicYearID;references:AcademicYearID"`
	lib.BaseModel                           /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (AcademicYear) TableName() string {
	return lib.GenerateTableName(lib.Academic, "academic_years")
}
