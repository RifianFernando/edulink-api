package controllers

import (
	"net/http"

	"github.com/edulink-api/database/models"
	"github.com/edulink-api/helper"
	"github.com/gin-gonic/gin"
)

// func GetAllStudentPersonalDataArchive(c *gin.Context) {
// 	// Get All all student personal data
// 	studentPersonalData, err := helper.GetAllStudentPersonalDataArchive()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Send the grouped result as a response
// 	c.JSON(http.StatusOK, gin.H{"student-personal-data": studentPersonalData})
// }

func GetAllStudentAttendanceArchive(c *gin.Context) {
	// get params
	academicYearStart := c.Param("academic_year_start")
	academicYearEnd := c.Param("academic_year_end")

	err := helper.ValidateAcademicYearInput(
		academicYearStart + "/" + academicYearEnd,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get All all student attendance
	studentAttendance, err := helper.GetAllStudentAttendanceArchive(
		academicYearStart,
		academicYearEnd,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Send the grouped result as a response
	c.JSON(http.StatusOK, gin.H{"student-attendance": studentAttendance})
}

func GetAllStudentScoreArchive(c *gin.Context) {
	// get params
	academicYearStart := c.Param("academic_year_start")
	academicYearEnd := c.Param("academic_year_end")
	academicSemesterYear := academicYearStart + "/" + academicYearEnd

	err := helper.ValidateAcademicYearInput(academicSemesterYear)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Search academic year ID
	var academicYear models.AcademicYear
	academicYear.AcademicYear = academicSemesterYear

	err = academicYear.GetAcademicYearByModel()
	if err != nil || academicYear.AcademicYearID == 0 {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": "academic year not found",
			},
		)
		return
	}

	classID := c.Param("class_id")

	// Get All all student score
	studentScore, err := helper.GetAllStudentScoreArchive(academicYear, classID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Send the grouped result as a response
	c.JSON(http.StatusOK, gin.H{"student-score": studentScore})
}

// func GetAllClassArchive(c *gin.Context) {
// 	// Get All all class
// 	class, err := helper.GetAllClassArchive()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Send the grouped result as a response
// 	c.JSON(http.StatusOK, gin.H{"class": class})
// }

// func GetAllScheduleArchive(c *gin.Context) {
// 	// Get All all schedule
// 	schedule, err := helper.GetAllScheduleArchive()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Send the grouped result as a response
// 	c.JSON(http.StatusOK, gin.H{"schedule": schedule})
// }

// func GetAllCalendarArchive(c *gin.Context) {
// 	// Get All all calendar
// 	calendar, err := helper.GetAllCalendarArchive()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Send the grouped result as a response
// 	c.JSON(http.StatusOK, gin.H{"calendar": calendar})
// }
