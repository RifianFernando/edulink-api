package models

import (
	"fmt"
	"time"

	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/database/migration/lib"
	"gorm.io/gorm"
)

type Session struct {
	SessionID     int64     `gorm:"primaryKey"`     // Unique session identifier
	UserID        int64     `gorm:"not null;index"` // References the User ID
	RefreshToken  string    `gorm:"not null"`       // Token for refreshing the session
	IPAddress     string    `gorm:"not null"`       // IP address of the user (optional)
	UserAgent     string    `gorm:"not null"`       // User agent string for the device/browser used
	ExpiresAt     time.Time `gorm:"not null"`       // Session expiration timestamp
	lib.BaseModel           // Includes CreatedAt, UpdatedAt, DeletedAt
}

func (Session) TableName() string {
	return lib.GenerateTableName(lib.Public, "sessions")
}

// is refresh token in the database by user id and refresh token
func (session *Session) GetSession() Session {
	// Get the session details from the database
	connections.DB.Where(&session).First(&session)

	return *session
}

// Check if the session exists in the database by user id and refresh token
func (session *Session) SessionExists() (bool, error) {
	result := connections.DB.Where(&session).First(&session)

	fmt.Println("session:", session)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, result.Error
	}

	return true, nil
}

// Insert a new session into the database
func (session *Session) InsertSession() error {
	// Insert the session details into the database
	result := connections.DB.Create(&session)

	// Check for errors
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Update an existing session in the database
func (session *Session) UpdateSession() error {
	// Update the session details in the database
	result := connections.DB.Save(&session)

	// Check for errors
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// bulk delete session
func (session *Session) DeleteSession() error {
	result := connections.DB.Unscoped().Delete(&session)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
