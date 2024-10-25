package migration

import (
	"time"

	"github.com/edulink-api/database/migration/lib"
)

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type Session struct {
	SessionID    int64    `gorm:"primaryKey;autoIncrement"` // Unique session identifier
	UserID       int64     `gorm:"not null;index;unique"`
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
