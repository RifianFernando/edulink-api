package helper

import (
	"net/http"
	"time"

	"github.com/edulink-api/database/models"
	request "github.com/edulink-api/request/attendance"
	"github.com/edulink-api/res"
	"github.com/gin-gonic/gin"
)

func HandleCreateUpdateStudentAttendance(
	c *gin.Context,
	request request.AllAttendanceRequest,
) (
	string,
	time.Time,
	[]models.ClassDateAttendanceStudent,
) {
	var allErrors []map[string]string

	// Bind the request JSON to the CreateStudentRequest struct
	if err := res.ResponseMessage(c.ShouldBindJSON(&request)); len(err) > 0 {
		allErrors = append(allErrors, err...)
	}

	// Validate the request
	if err := request.ValidateAllAttendance(); len(err) > 0 {
		allErrors = append(allErrors, err...)
	}

	if len(allErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": allErrors,
		})
		c.Abort()
	}

	var (
		err     error
		ClassID string
		Date    time.Time
	)

	ClassID, Date, err = GetHomeRoomTeacherByTeacherID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
	}

	// get attendance
	var attendances []models.ClassDateAttendanceStudent
	for _, attendance := range request.AttendanceRequest {
		attendances = append(attendances, models.ClassDateAttendanceStudent{
			StudentID: attendance.StudentID,
			Reason:    attendance.Reason,
		})
	}

	return ClassID, Date, attendances
}
