package models

import (
	"fmt"
	"strings"

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
	User User `gorm:"foreignKey:UserID;references:UserID"` // Belongs-to with User
}

func (StaffModel) TableName() string {
	return lib.GenerateTableName(lib.Administration, "staffs")
}

// Get staff by model
func (staff *StaffModel) GetStaffByModel() error {
	if err := connections.DB.Where(&staff).Preload("User").First(&staff).Error; err != nil {
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

func (staff *Staff) UpdateStaffByModel() error {
	result := connections.DB.Model(&staff).Updates(&staff)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("Staff not found")
	}

	return nil
}

func (staff *StaffModel) UpdateStaffByID() error {
	tx := connections.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	result := tx.Model(&staff.Staff).Updates(
		&Staff{
			Position: staff.Position,
		},
	)
	if result.Error != nil {
		tx.Rollback()
		// return result.Error
		return fmt.Errorf("error staff: %v", result.Error)
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("staff not found")
	}

	userData := staff.User
	result = tx.Model(&staff.User).
		Where("user_id = ?", staff.UserID).
		Updates(
			&User{
				UserID:           userData.UserID,
				UserName:         userData.UserName,
				UserGender:       userData.UserGender,
				UserPlaceOfBirth: userData.UserPlaceOfBirth,
				UserDateOfBirth:  userData.UserDateOfBirth,
				UserReligion:     userData.UserReligion,
				UserAddress:      userData.UserAddress,
				UserPhoneNum:     userData.UserPhoneNum,
				UserEmail:        userData.UserEmail,
			},
		)
	if result.Error != nil {
		tx.Rollback()
		// return result.Error
		return fmt.Errorf("error user: %v", result.Error)
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return tx.Commit().Error
}

func (staff *Staff) DeleteStaffByID() error {
	result := connections.DB.Delete(&staff)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func CreateAllStaffs(staffs []StaffModel) error {
	tx := connections.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, staffModel := range staffs {
		if err := createUserAndStaff(tx, staffModel); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func createUserAndStaff(tx *gorm.DB, staffModel StaffModel) error {
	user := createUserFromModel(staffModel.User)
	userResult, err := createUser(tx, user)
	if err != nil {
		return err
	}

	staff := Staff{
		UserID:   userResult.UserID,
		Position: staffModel.Position,
	}
	if err := createStaff(tx, staff); err != nil {
		return err
	}

	return nil
}

func createUserFromModel(userModel User) User {
	return User{
		UserName:         userModel.UserName,
		UserGender:       userModel.UserGender,
		UserPlaceOfBirth: userModel.UserPlaceOfBirth,
		UserDateOfBirth:  userModel.UserDateOfBirth,
		UserReligion:     userModel.UserReligion,
		UserAddress:      userModel.UserAddress,
		UserPhoneNum:     userModel.UserPhoneNum,
		UserEmail:        userModel.UserEmail,
	}
}

func createUser(tx *gorm.DB, user User) (User, error) {
	result := tx.Create(&user)
	if result.Error != nil {
		if pgErr, ok := result.Error.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				if strings.Contains(pgErr.ConstraintName, "phone") {
					return User{}, fmt.Errorf("staff with phone number %s already exists", user.UserPhoneNum)
				} else if strings.Contains(pgErr.ConstraintName, "email") {
					return User{}, fmt.Errorf("staff with email %s already exists", user.UserEmail)
				}
			}
		}
		return User{}, result.Error
	}
	return user, nil
}

func createStaff(tx *gorm.DB, staff Staff) error {
	result := tx.Create(&staff)
	if result.Error != nil {
		return result.Error
	}
	// 	return
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"staffs": staffs,
	// 	})
	// }
	return nil
}
