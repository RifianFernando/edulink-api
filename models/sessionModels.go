package models

import (
	"time"

	"github.com/skripsi-be/database/migration/lib"
)

type Session struct {
	SessionID     int64    `gorm:"primaryKey"` // Unique session identifier
	UserID        int64     `gorm:"not null;index"`      // References the User ID
	SessionToken  string    `gorm:"not null"`            // Token for authentication (JWT or random string)
	RefreshToken  string    `gorm:"not null"`            // Token for refreshing the session
	IPAddress     string    `gorm:"not null"`            // IP address of the user (optional)
	UserAgent     string    `gorm:"not null"`            // User agent string for the device/browser used
	ExpiresAt     time.Time `gorm:"not null"`            // Session expiration timestamp
	lib.BaseModel           // Includes CreatedAt, UpdatedAt, DeletedAt
}

func (Session) TableName() string {
	return lib.GenerateTableName(lib.Public, "sessions")
}
