package migration

import (
	"time"
)

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

func (g *Gender) Scan(value interface{}) error {
	*g = Gender(value.([]byte))
	return nil
}

/* 
	* see the documentation here
	* https://gorm.io/docs/data_types.html
*/
type User struct {
	UserID            int64  `gorm:"primaryKey;autoIncrement"`
	UserName          string `gorm:"unique;not null"`
	UserGender        Gender `gorm:"type:VARCHAR(6);not null"`
	UserPlaceOfBirth  string `gorm:"not null"`
	UserDateOfBirth   time.Time `gorm:"not null"`
	UserAddress       string `gorm:"not null"`
	UserNumPhone      string `gorm:"unique;not null"`
	UserEmail         string `gorm:"unique;not null"`
	UserPassword      string `gorm:"not null"`
	Teacher           Teacher `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Staff             Staff `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	BaseModel /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
	* see the documentation here about conventions
	* https://gorm.io/docs/conventions.html
*/
func (User) TableName() string {
	return GenerateTableName(Public, "users")
}
