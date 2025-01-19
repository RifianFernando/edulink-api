package helper

import (
	"errors"

	"github.com/edulink-api/database/models"
	request "github.com/edulink-api/request/staff"
	"github.com/gin-gonic/gin"
)

func PrepareStaffs(requestedStaffs []request.InsertStaffRequest, c *gin.Context) ([]models.StaffModel, error) {
	numPhoneMap := make(map[string]bool)
	emailMap := make(map[string]bool)

	var staffs []models.StaffModel

	for index, staff := range requestedStaffs {
		if err := checkForDuplicatesStaff(
			staff,
			index,
			numPhoneMap,
			emailMap,
		); err != nil {
			return nil, err
		}

		DateOfBirth, err := staff.ParseDates()
		if err != nil {
			return nil, err
		}

		// Add the validated staff to the slice
		staffs = append(staffs, models.StaffModel{
			Staff: models.Staff{
				Position: staff.Position,
			},
			User: models.User{
				UserName:         staff.UserName,
				UserGender:       staff.UserGender,
				UserPlaceOfBirth: staff.UserPlaceOfBirth,
				UserDateOfBirth:  DateOfBirth,
				UserReligion:     staff.UserReligion,
				UserAddress:      staff.UserAddress,
				UserPhoneNum:     staff.UserPhoneNum,
				UserEmail:        staff.UserEmail,
			},
		})
		index++
	}

	return staffs, nil
}

func checkForDuplicatesStaff(
	staff request.InsertStaffRequest,
	index int,
	numPhoneMap,
	emailMap map[string]bool,
) error {
	if numPhoneMap[staff.UserPhoneNum] {
		return errors.New(CustomErrorForDuplicate("Staff Phone Number", staff.UserPhoneNum, index))
	}
	if emailMap[staff.UserEmail] {
		return errors.New(CustomErrorForDuplicate("Staff Email", staff.UserEmail, index))
	}

	emailMap[staff.UserEmail] = true
	numPhoneMap[staff.UserPhoneNum] = true

	return nil
}
