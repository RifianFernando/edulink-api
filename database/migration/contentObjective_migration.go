package migration

/* 
	* see the documentation here
	* https://gorm.io/docs/data_types.html
*/
type ContentObjective struct {
	SyllabusID int64 `gorm:"not null"`
	ContentObjectiveContentTopic string `gorm:"not null"`
	ContentObjectiveLearningObjective string `gorm:"not null"`
	ContentObjectiveDuration int32 `gorm:"not null"`
	BaseModel /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
	* see the documentation here about conventions
	* https://gorm.io/docs/conventions.html
*/
func (ContentObjective) TableName() string {
	return GenerateTableName(Academic, "content_objectives")
}
