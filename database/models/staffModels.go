package models

import (
	"fmt"

	"github.com/edulink-api/connections"
	"github.com/edulink-api/database/migration/lib"
	"gorm.io/gorm"
)

type Staff struct {
	StaffId  int64  `gorm:"primaryKey"`
	UserID   int64  `json:"id_user" binding:"required"`
	Position string `json:"user_position" binding:"required"`
	lib.BaseModel
}

func (Staff) TableName() string {
	return lib.GenerateTableName(lib.Administration, "staffs")
}

// Get staff by model
func (staff *Staff) GetStaffByModel() error {
	if err := connections.DB.Where(&staff).First(&staff).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Staff not found")
		}
		return err
	}

	return nil
}
