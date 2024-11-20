package migration

import (
	"time"

	"github.com/edulink-api/database/migration/lib"
)

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

type Religion string

const (
	Islam     Religion = "Islam"
	Christian Religion = "Kristen Protestan"
	Catholic  Religion = "Kristen Katolik"
	Hindu     Religion = "Hindu"
	Buddha    Religion = "Buddha"
	Konghucu  Religion = "Konghucu"
)

func (g *Gender) Scan(value interface{}) error {
	*g = Gender(value.([]byte))
	return nil
}

func (r *Religion) Scan(value interface{}) error {
	*r = Religion(value.([]byte))
	return nil
}

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type User struct {
	UserID           int64     `gorm:"primaryKey;autoIncrement"`
	UserName         string    `gorm:"not null"`
	UserGender       Gender    `gorm:"type:VARCHAR(6);not null"`
	UserPlaceOfBirth string    `gorm:"not null"`
	UserDateOfBirth  time.Time `gorm:"not null"`
	UserReligion     Religion  `gorm:"type:VARCHAR(17);not null"`
	UserAddress      string    `gorm:"not null;type:VARCHAR(200)"`
	UserNumPhone     string    `gorm:"unique;not null"`
	UserEmail        string    `gorm:"unique;not null"`
	UserPassword     string    `gorm:"not null"`
	Teacher          Teacher   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Staff            Staff     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Admin            Admin     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Session          Session   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	lib.BaseModel              /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (User) TableName() string {
	return lib.GenerateTableName(lib.Public, "users")
}
