package models

import (
	"fmt"
	"time"

	"github.com/edulink-api/connections"
	"github.com/edulink-api/database/migration/lib"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Teacher struct {
	TeacherID    int64 `gorm:"primaryKey"`
	UserID       int64 `json:"id_user" binding:"required"`
	TeachingHour int32 `json:"teaching_hour" bainding:"required"`
	lib.BaseModel
}

type TeacherModel struct {
	Teacher
	ClassNames     []ClassName      `gorm:"foreignKey:TeacherID;references:TeacherID"`
	TeacherSubject []TeacherSubject `gorm:"foreignKey:TeacherID;references:TeacherID"`
	User           User             `gorm:"foreignKey:UserID;references:UserID"` // Belongs-to with User
	// Scores       []Score     `gorm:"foreignKey:TeacherID;references:TeacherID;constraint:OnUpdate:SET NULL,OnDelete:SET NULL"`
}

func (Teacher) TableName() string {
	return lib.GenerateTableName(lib.Academic, "teachers")
}

// Create teacher
func (teacher *Teacher) CreateTeacher() error {
	result := connections.DB.Create(&teacher)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Get all teacher
func (teacher *TeacherModel) GetAllUserTeachersWithUser() (
	teachers []TeacherModel,
	msg string,
) {
	result := connections.DB.Preload("User").Preload("TeacherSubject.Subject").Find(&teachers)
	if result.Error != nil {
		return nil, result.Error.Error()
	} else if result.RowsAffected == 0 {
		return nil, "No user teacher found"
	}

	return teachers, ""
}

type GetTeacherByIDWithoutPassword struct {
	TeacherID        int64     `json:"teacher_id"`
	UserID           int64     `json:"user_id"`
	UserName         string    `json:"name" binding:"required"`
	UserGender       string    `json:"gender" binding:"required,oneof=Male Female"`
	UserPlaceOfBirth string    `json:"place_of_birth" binding:"required"`
	UserDateOfBirth  time.Time `json:"date_of_birth"`
	UserReligion     string    `json:"religion" binding:"required" validate:"required,oneof='Islam' 'Kristen Protestan' 'Kristen Katolik' 'Hindu' 'Buddha' 'Konghucu'"`
	UserAddress      string    `json:"address" binding:"required" validate:"required,min=10,max=200"`
	UserPhoneNum     string    `json:"num_phone" binding:"required,e164"`
	UserEmail        string    `json:"email" binding:"required,email"`
	TeachingHour     int32     `json:"teaching_hour"`
	Subject          []Subject `json:"subject"`
}

// Get teacher by id
func (teacher *TeacherModel) GetTeacherById(id string) (GetTeacherByIDWithoutPassword, error) {

	result := connections.DB.Preload("User").Preload("TeacherSubject.Subject").First(&teacher, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return GetTeacherByIDWithoutPassword{}, fmt.Errorf("no teacher found")
		}
		return GetTeacherByIDWithoutPassword{}, result.Error
	}

	var teacherDTO GetTeacherByIDWithoutPassword

	teacherDTO.TeacherID = teacher.TeacherID
	teacherDTO.UserID = teacher.UserID
	teacherDTO.UserName = teacher.User.UserName
	teacherDTO.UserGender = teacher.User.UserGender
	teacherDTO.UserPlaceOfBirth = teacher.User.UserPlaceOfBirth
	teacherDTO.UserDateOfBirth = teacher.User.UserDateOfBirth
	teacherDTO.UserReligion = teacher.User.UserReligion
	teacherDTO.UserAddress = teacher.User.UserAddress
	teacherDTO.UserPhoneNum = teacher.User.UserPhoneNum
	teacherDTO.UserEmail = teacher.User.UserEmail
	teacherDTO.TeachingHour = teacher.TeachingHour
	for _, subject := range teacher.TeacherSubject {
		teacherDTO.Subject = append(teacherDTO.Subject, subject.Subject)
	}

	return teacherDTO, nil
}

// Get teacher by user id
func (teacher *Teacher) GetTeacherByModel() error {
	result := connections.DB.Where(&teacher).First(&teacher)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return fmt.Errorf("teacher not found")
		}
		return result.Error
	}

	return nil
}

// Update teacher and user
func (teacher *TeacherModel) UpdateTeacherById(teacherData *TeacherModel) error {
	tx := connections.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	result := tx.Model(&teacher.Teacher).Updates(
		&Teacher{
			TeachingHour: teacherData.TeachingHour,
		},
	)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("teacher not found")
	}

	userData := teacherData.User
	result = tx.Model(&teacher.User).Updates(
		&User{
			UserName:         userData.UserName,
			UserGender:       userData.UserGender,
			UserPlaceOfBirth: userData.UserPlaceOfBirth,
			UserDateOfBirth:  userData.UserDateOfBirth,
			UserAddress:      userData.UserAddress,
			UserPhoneNum:     userData.UserPhoneNum,
			UserReligion:     userData.UserReligion,
			UserEmail:        userData.UserEmail,
		},
	)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	// delete teacher subject
	result = tx.Where("teacher_id = ?", teacherData.TeacherID).Delete(&TeacherSubject{})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	// upsert teacher subject if found deleted at will null or restored
	for _, teacherSubject := range teacherData.TeacherSubject {
		result = tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "teacher_id"}, {Name: "subject_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"deleted_at"}),
		}).Create(&teacherSubject)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}

	return tx.Commit().Error
}

// Delete teacher by id
func (teacher *Teacher) DeleteTeacherById() error {
	result := connections.DB.Delete(&teacher)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// var className []ClassNameModel
type Output struct {
	SubjectID   int    `json:"subject_id"`
	ClassNameID int    `json:"class_name_id"`
	Grade       string `json:"grade"`
	Name        string `json:"name"`
	SubjectName string `json:"subject_name"`
}

var output []Output

func GetTeacherTeachingClassList(teacherID string) ([]Output, error) {
	query := `
	select 
		s.subject_id,
		cn.class_name_id,
		g.grade ,
		cn."name" ,
		s.subject_name
	from academic.teachers t 
	join academic.teacher_subjects ts on t.teacher_id  = ts.teacher_id 
	join academic.teaching_class_subjects tcs on ts.teacher_subject_id = tcs.teacher_subject_id 
	join academic.class_names cn on tcs.class_name_id = cn.class_name_id 
	join academic.grades g on cn.grade_id = g.grade_id 
	join academic.subjects s on ts.subject_id = s.subject_id 
	where t.teacher_id = ?`

	result := connections.DB.Raw(query, teacherID).Scan(&output)

	if result.Error != nil {
		return nil, result.Error
	}

	return output, nil
}
