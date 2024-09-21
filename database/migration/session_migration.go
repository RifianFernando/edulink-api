package migration

import (
	"time"

	"github.com/skripsi-be/database/migration/lib"
)

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type Session struct {
	SessionID    string    `gorm:"primaryKey;size:255"`
	UserID       int64     `gorm:"not null;index"`
	SessionToken string    `gorm:"not null"`
	RefreshToken string    `gorm:"not null"`
	IPAddress    string    `gorm:"not null"`
	UserAgent    string    `gorm:"not null"`
	ExpiresAt    time.Time `gorm:"not null"`
	lib.BaseModel
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (Session) TableName() string {
	return lib.GenerateTableName(lib.Public, "sessions")
}
