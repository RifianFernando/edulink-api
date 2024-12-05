package helper

import (
	"fmt"
	"strconv"
	"time"

	"github.com/edulink-api/database/models"
	"github.com/edulink-api/res"
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
			res.AbortUnauthorized(c)
			return "", time.Time{}, err
		}

		// get class id
		var className models.ClassName
		className.TeacherID = teacher.TeacherID
		classes, err := className.GetHomeRoomTeacherByTeacherID()
		if err != nil {
			res.AbortUnauthorized(c)
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
			res.AbortUnauthorized(c)
			return "", time.Time{}, fmt.Errorf("forbidden")
		}
	}

	return ClassID, Date, nil
}
