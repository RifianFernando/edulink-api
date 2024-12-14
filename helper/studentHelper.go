package helper

import (
	"errors"
	"strconv"
	"time"

	"github.com/edulink-api/database/models"
	"github.com/edulink-api/request/student"
	"github.com/gin-gonic/gin"
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

		// make the server load is heavy should not check the existence of the student one by one but create a query to check all the existence at once
		// if err := checkStudentExistence(student, index); err != nil {
		// 	return nil, err
		// }

		// this make server load heavy too
		// classIDStr := strconv.FormatInt(student.ClassNameID, 10)
		// if err := checkClassExistence(classIDStr, index); err != nil {
		// 	return nil, err
		// }

		// Add the validated student to the slice
		students = append(students, models.Student{
			ClassNameID:           student.ClassNameID,
			StudentName:           student.StudentName,
			StudentNISN:           student.StudentNISN,
			StudentGender:         student.StudentGender,
			StudentPlaceOfBirth:   student.StudentPlaceOfBirth,
			StudentDateOfBirth:    parsedDates.DateOfBirth,
			StudentReligion:       student.StudentReligion,
			StudentAddress:        student.StudentAddress,
			StudentPhoneNumber:       student.StudentPhoneNumber,
			StudentEmail:          student.StudentEmail,
			StudentAcceptedDate:   parsedDates.AcceptedDate,
			StudentSchoolOfOrigin:   student.StudentSchoolOfOrigin,
			StudentFatherName:     student.StudentFatherName,
			StudentFatherJob:      student.StudentFatherJob,
			StudentFatherPhoneNumber: student.StudentFatherPhoneNumber,
			StudentMotherName:     student.StudentMotherName,
			StudentMotherJob:      student.StudentMotherJob,
			StudentMotherPhoneNumber: student.StudentMotherPhoneNumber,
		})
	}

	return students, nil
}

func customErrorForDuplicate(property string, atribute string, index int) string {
	return "Duplicate " + property + ": " + atribute + " on index: " + strconv.Itoa(index + 1)
}

func checkForDuplicates(student request.InsertStudentRequest, index int, nameMap, nisnMap, numPhoneMap, emailMap map[string]bool) error {
	// name can be duplicate
	// if nameMap[student.StudentName] {
	// 	return errors.New(customErrorForDuplicate("StudentName", student.StudentName, index))
	// }
	if nisnMap[student.StudentNISN] {
		return errors.New(customErrorForDuplicate("StudentNISN", student.StudentNISN, index))
	}
	if numPhoneMap[student.StudentPhoneNumber] {
		return errors.New(customErrorForDuplicate("StudentPhoneNumber", student.StudentPhoneNumber, index))
	}
	if emailMap[student.StudentEmail] {
		return errors.New(customErrorForDuplicate("StudentEmail", student.StudentEmail, index))
	}

	nameMap[student.StudentName] = true
	nisnMap[student.StudentNISN] = true
	numPhoneMap[student.StudentPhoneNumber] = true
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
		{student.StudentPhoneNumber, "number phone"},
		{student.StudentEmail, "email"},
	}

	for _, criteria := range searchCriteria {
		var studentSearch = models.Student{}
		switch criteria.field {
		// student name can be duplicate
		// case "name":
		// 	studentSearch.StudentName = criteria.value
		case "NISN":
			studentSearch.StudentNISN = criteria.value
		case "number phone":
			studentSearch.StudentPhoneNumber = criteria.value
		case "email":
			studentSearch.StudentEmail = criteria.value
		}

		result, err := studentSearch.GetStudent()
		if err != nil {
			return err
		}

		if result.StudentID != 0 || result.ClassNameID != 0 {
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
