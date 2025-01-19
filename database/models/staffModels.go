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

type StaffModel struct {
	Staff
	User User `gorm:"foreignKey:UserID;references:UserID"` // Belongs-to with User
	// Scores       []Score     `gorm:"foreignKey:TeacherID;references:TeacherID;constraint:OnUpdate:SET NULL,OnDelete:SET NULL"`
}

func (StaffModel) TableName() string {
	return lib.GenerateTableName(lib.Administration, "staffs")
}

// Get staff by model
func (staff *StaffModel) GetStaffByModel() error {
	if err := connections.DB.Where(&staff).First(&staff).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Staff not found")
		}
		return err
	}

	return nil
}

func (staff *Staff) CreateStaff() error {
	result := connections.DB.Create(&staff)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (staff *StaffModel) GetAllUserStaffWithUser() (
	staffs []StaffModel,
	msg string,
) {
	result := connections.DB.Preload("User").Find(&staffs)
	if result.Error != nil {
		return nil, result.Error.Error()
	} else if result.RowsAffected == 0 {
		return nil, "No user staff found"
	}

	return staffs, ""
}
