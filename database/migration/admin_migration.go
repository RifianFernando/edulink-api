package migration

import "github.com/skripsi-be/database/migration/lib"

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type Admin struct {
	AdminID       int64  `gorm:"primaryKey;autoIncrement"`
	UserID        int64  `gorm:"not null"`
	Position      string `gorm:"not null"`
	lib.BaseModel        /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (Admin) TableName() string {
	return lib.GenerateTableName(lib.Public, "admins")
}
