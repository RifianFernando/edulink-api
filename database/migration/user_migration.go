package migration

import (
	"time"
)

type User struct {
	UserID            int64  `gorm:"primaryKey;autoIncrement"`
	UserName          string `gorm:"unique;not null"`
	UserGender        string `gorm:"unique;not null"`
	UserPlaceOfBirth  string `gorm:"not null"`
	UserDateOfBirth   time.Time `gorm:"not null"`
	UserAddress       string `gorm:"not null"`
	UserNumPhone      string `gorm:"unique;not null"`
	UserEmail         string `gorm:"unique;not null"`
	UserPassword      string `gorm:"not null"`
	Teacher           Teacher `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Staff             Staff `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
