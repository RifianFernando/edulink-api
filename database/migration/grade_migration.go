package migration

/* 
	* see the documentation here
	* https://gorm.io/docs/data_types.html
*/
type Grade struct {
	StudentID  int64 `gorm:"not null"`
	AssignmentID int64 `gorm:"not null"`
	TeacherID  int64 `gorm:"not null"`
	SubjectID  int64 `gorm:"not null"`
	Grade      string `gorm:"not null"`
	BaseModel /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
	* see the documentation here about conventions
	* https://gorm.io/docs/conventions.html
*/
func (Grade) TableName() string {
	return GenerateTableName(Academic, "grades")
}
