package migration

/* 
	* see the documentation here
	* https://gorm.io/docs/data_types.html
*/
type DomainAchievement struct {
	SyllabusID int64 `gorm:"not null"`
	DomainAchievementLearningAchievement string `gorm:"not null"`
	BaseModel /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
	DomainAchievementDomain string `gorm:"not null"`
}

/*
	* see the documentation here about conventions
	* https://gorm.io/docs/conventions.html
*/
func (DomainAchievement) TableName() string {
	return GenerateTableName(Academic, "domain_achievements")
}
