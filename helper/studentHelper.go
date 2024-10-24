package helper

import (
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/models"
	"github.com/skripsi-be/request"
)

func PrepareStudents(requestedStudents []request.InsertStudentRequest, c *gin.Context) ([]models.Student, error) {
	nameMap := make(map[string]bool)
	nisnMap := make(map[string]bool)
	numPhoneMap := make(map[string]bool)
	emailMap := make(map[string]bool)

	var students []models.Student

	for index, student := range requestedStudents {
		index++
		if err := checkForDuplicates(student, index, nameMap, nisnMap, numPhoneMap, emailMap); err != nil {
			return nil, err
		}

		parsedDates, err := parseStudentDates(student, index)
		if err != nil {
			return nil, err
		}

		if err := checkStudentExistence(student, index); err != nil {
			return nil, err
		}

		classIDStr := strconv.FormatInt(student.ClassID, 10)
		if err := checkClassExistence(classIDStr, index); err != nil {
			return nil, err
		}

		// Add the validated student to the slice
		students = append(students, models.Student{
			ClassID:               student.ClassID,
			StudentName:           student.StudentName,
			StudentNISN:           student.StudentNISN,
			StudentGender:         student.StudentGender,
			StudentPlaceOfBirth:   student.StudentPlaceOfBirth,
			StudentDateOfBirth:    parsedDates.DateOfBirth,
			StudentReligion:       student.StudentReligion,
			StudentAddress:        student.StudentAddress,
			StudentNumPhone:       student.StudentNumPhone,
			StudentEmail:          student.StudentEmail,
			StudentAcceptedDate:   parsedDates.AcceptedDate,
			StudentSchoolOrigin:   student.StudentSchoolOrigin,
			StudentFatherName:     student.StudentFatherName,
			StudentFatherJob:      student.StudentFatherJob,
			StudentFatherNumPhone: student.StudentFatherNumPhone,
			StudentMotherName:     student.StudentMotherName,
			StudentMotherJob:      student.StudentMotherJob,
			StudentMotherNumPhone: student.StudentMotherNumPhone,
		})
	}

	return students, nil
}

func checkForDuplicates(student request.InsertStudentRequest, index int, nameMap, nisnMap, numPhoneMap, emailMap map[string]bool) error {
	if nameMap[student.StudentName] {
		return errors.New("Duplicate StudentName: " + student.StudentName + " on index: " + strconv.Itoa(index))
	}
	if nisnMap[student.StudentNISN] {
		return errors.New("Duplicate StudentNISN: " + student.StudentNISN + " on index: " + strconv.Itoa(index))
	}
	if numPhoneMap[student.StudentNumPhone] {
		return errors.New("Duplicate StudentNumPhone: " + student.StudentNumPhone + " on index: " + strconv.Itoa(index))
	}
	if emailMap[student.StudentEmail] {
		return errors.New("Duplicate StudentEmail: " + student.StudentEmail + " on index: " + strconv.Itoa(index))
	}

	nameMap[student.StudentName] = true
	nisnMap[student.StudentNISN] = true
	numPhoneMap[student.StudentNumPhone] = true
	emailMap[student.StudentEmail] = true
	return nil
}

func parseStudentDates(student request.InsertStudentRequest, index int) (parsedDates struct {
	DateOfBirth  time.Time
	AcceptedDate time.Time
}, err error) {
	parsedDates.DateOfBirth, err = time.Parse("2006-01-02", student.DateOfBirth)
	if err != nil {
		return parsedDates, errors.New("Invalid date format on index: " + strconv.Itoa(index))
	}

	parsedDates.AcceptedDate, err = time.Parse("2006-01-02", student.AcceptedDate)
	if err != nil {
		return parsedDates, errors.New("Invalid date format on index: " + strconv.Itoa(index))
	}

	return parsedDates, nil
}

func checkStudentExistence(student request.InsertStudentRequest, index int) error {
	searchCriteria := []struct {
		value string
		field string
	}{
		{student.StudentName, "name"},
		{student.StudentNISN, "NISN"},
		{student.StudentNumPhone, "number phone"},
		{student.StudentEmail, "email"},
	}

	for _, criteria := range searchCriteria {
		var studentSearch = models.Student{}
		switch criteria.field {
		case "name":
			studentSearch.StudentName = criteria.value
		case "NISN":
			studentSearch.StudentNISN = criteria.value
		case "number phone":
			studentSearch.StudentNumPhone = criteria.value
		case "email":
			studentSearch.StudentEmail = criteria.value
		}

		result, err := studentSearch.GetStudent()
		if err != nil {
			return err
		}

		if result.StudentID != 0 || result.ClassID != 0 {
			return errors.New("Student " + criteria.field + ": " + criteria.value + " already exist on index: " + strconv.Itoa(index))
		}
	}

	return nil
}

func checkClassExistence(classID string, index int) error {
	var class models.ClassName
	resultClass, err := class.GetClassNameById(classID)
	if err != nil {
		return err
	}
	if resultClass.ClassNameID == 0 {
		return errors.New("Class with id: " + classID + " doesn't exist on index: " + strconv.Itoa(index))
	}
	return nil
}
