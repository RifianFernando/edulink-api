package models

import (
	"fmt"
	"time"

	"github.com/edulink-api/connections"
	"github.com/edulink-api/database/migration/lib"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type Staff struct {
	StaffID  int64  `gorm:"primaryKey"`
	UserID   int64  `json:"id_user" binding:"required"`
	Position string `json:"user_position" binding:"required"`
	lib.BaseModel
}

func (Staff) TableName() string {
	return lib.GenerateTableName(lib.Administration, "staffs")
}

type StaffModel struct {
	Staff
	User User `gorm:"foreignKey:UserID;references:UserID"`
}

// get all staff

type GetStaffByIDWithoutPassword struct {
	StaffID          int64     `json:"staff_id"`
	UserID           int64     `json:"user_id"`
	UserName         string    `json:"name" binding:"required"`
	UserGender       string    `json:"gender" binding:"required,oneof=Male Female"`
	UserPlaceOfBirth string    `json:"place_of_birth" binding:"required"`
	UserDateOfBirth  time.Time `json:"date_of_birth"`
	UserReligion     string    `json:"religion" binding:"required" validate:"required,oneof='Islam' 'Kristen Protestan' 'Kristen Katolik' 'Hindu' 'Buddha' 'Konghucu'"`
	UserAddress      string    `json:"address" binding:"required" validate:"required,min=10,max=200"`
	UserPhoneNum     string    `json:"num_phone" binding:"required,e164"`
	UserEmail        string    `json:"email" binding:"required,email"`
	Position         string    `json:"position" binding:"required"`
}

func (staff *StaffModel) GetAllStaff() (staffDTO []GetStaffByIDWithoutPassword, err error) {
	var staffs []StaffModel
	if err = connections.DB.Preload("User").Find(&staffs).Error; err != nil {
		return nil, err
	}

	for _, staff := range staffs {
		staffDTO = append(staffDTO, GetStaffByIDWithoutPassword{
			StaffID:          staff.Staff.StaffID,
			UserID:           staff.Staff.StaffID,
			UserName:         staff.User.UserName,
			UserGender:       staff.User.UserGender,
			UserPlaceOfBirth: staff.User.UserPlaceOfBirth,
			UserDateOfBirth:  staff.User.UserDateOfBirth,
			UserReligion:     staff.User.UserReligion,
			UserAddress:      staff.User.UserAddress,
			UserPhoneNum:     staff.User.UserPhoneNum,
			UserEmail:        staff.User.UserEmail,
			Position:         staff.Staff.Position,
		})
	}

	return staffDTO, nil
}

// Get staff by model
func (staff *StaffModel) GetStaffByModel() (staffDTO GetStaffByIDWithoutPassword, err error) {
	if err := connections.DB.Preload("User").Where(&staff).First(&staff).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return staffDTO, fmt.Errorf("staff not found")
		}
		return staffDTO, err
	}

	staffDTO.StaffID = staff.Staff.StaffID
	staffDTO.UserID = staff.Staff.UserID
	staffDTO.UserName = staff.User.UserName
	staffDTO.UserGender = staff.User.UserGender
	staffDTO.UserPlaceOfBirth = staff.User.UserPlaceOfBirth
	staffDTO.UserDateOfBirth = staff.User.UserDateOfBirth
	staffDTO.UserReligion = staff.User.UserReligion
	staffDTO.UserAddress = staff.User.UserAddress
	staffDTO.UserPhoneNum = staff.User.UserPhoneNum
	staffDTO.UserEmail = staff.User.UserEmail
	staffDTO.Position = staff.Staff.Position

	return staffDTO, nil
}

// Create Staff
func (staff *Staff) CreateStaff() error {
	if err := connections.DB.Create(&staff).Error; err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23503" && pgErr.ConstraintName == "fk_public_users_staff" {
				return fmt.Errorf("user not found")
			} else if pgErr.Code == "23505" {
				return fmt.Errorf("staff already exist")
			}
		}
		return err
	}

	return nil
}
