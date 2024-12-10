package helper

import (
	"fmt"
	"strconv"
	"time"

	"github.com/edulink-api/database/models"
	"github.com/gin-gonic/gin"
)

func GetHomeRoomTeacherByTeacherID(c *gin.Context) (string, time.Time, error) {
	// get user role
	userRole, _ := c.Get("user_type")
	userID, _ := c.Get("user_id")

	Date, err := time.Parse("2006-01-02", c.Param("date"))
	// var student models.StudentModel
	ClassID := c.Param("class_id")
	if err != nil {
		Date = time.Now()
	}

	if userRole != "admin" && userRole != "staff" {
		// check homeroom teacher class that he is assigned to
		var teacher models.Teacher
		teacher.UserID = userID.(int64)
		err := teacher.GetTeacherByModel()
		if err != nil {
			return "", time.Time{}, err
		}

		// get class id
		var className models.ClassName
		className.TeacherID = teacher.TeacherID
		classes, err := className.GetHomeRoomTeacherByTeacherID()
		if err != nil {
			return "", time.Time{}, err
		}

		// check if there's any class that the teacher is assigned to
		var isAssigned bool
		for _, className := range classes {
			if ClassID == strconv.FormatInt(className.ClassNameID, 10) {
				isAssigned = true
				break
			}
		}

		if !isAssigned {
			return "", time.Time{}, fmt.Errorf("forbidden")
		}
	}

	return ClassID, Date, nil
}

// TODO: use this function in controllers/scoringController.go
func IsTeachingClassSubjectExist(userID any, subjectID string, classNameID string) (models.Teacher, error) {

	// get teacher id
	var teacher models.Teacher
	teacher.UserID = userID.(int64)
	err := teacher.GetTeacherByModel()
	if err != nil || teacher.TeacherID == 0 {
		return teacher, fmt.Errorf("teacher not found")
	}
	teacherID := strconv.FormatInt(teacher.TeacherID, 10)
	classNameIDParsed, err := strconv.ParseInt(classNameID, 10, 64)	
	if err != nil {
		return teacher, fmt.Errorf("invalid class name id")
	}
	var result []models.TeacherSubjectGrade
	result, err = models.GetTeachingSubjectBySubjectID(
		subjectID,
		teacherID,
	)
	if err != nil {
		return teacher, err
	}

	var isExist = false
	for _, teacher := range result {
		if len(teacher.TeachingClassSubject) > 0 {
			for _, classSubject := range teacher.TeachingClassSubject {
				if classSubject.ClassNameID == classNameIDParsed {
					isExist = true
					break
				}
			}
		}
	}
	if !isExist {
		return teacher, fmt.Errorf("teacher is not teaching this class")
	}

	return teacher, nil
}
