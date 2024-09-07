package migration

import "github.com/skripsi-be/database/migration/lib"

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
 */
type Syllabus struct {
	SyllabusID         int64               `gorm:"primaryKey;autoIncrement"`
	SubjectID          int64               `gorm:"not null"`
	ContentObjectives  []ContentObjective  `gorm:"foreignKey:SyllabusID;references:SyllabusID"`
	DomainAchievements []DomainAchievement `gorm:"foreignKey:SyllabusID;references:SyllabusID"`
	SyllabusDetails    []SyllabusDetail    `gorm:"foreignKey:SyllabusID;references:SyllabusID"`
	lib.BaseModel                          /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (Syllabus) TableName() string {
	return lib.GenerateTableName(lib.Academic, "syllabuses")
}
