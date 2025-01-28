package controllers

import (
	"net/http"

	"github.com/edulink-api/database/models"
	request "github.com/edulink-api/request/event"
	"github.com/edulink-api/res"
	"github.com/gin-gonic/gin"
)

func CreateEvent(c *gin.Context) {
	var req request.InsertEventScheduleRequest
	var allErrors []map[string]string

	// Bind the request JSON to the CreateStudentRequest struct
	if err := res.ResponseMessage(c.ShouldBindJSON(&req)); len(err) > 0 {
		allErrors = append(allErrors, err...)
	}

	// Validate the request
	if err := req.Validate(); len(err) > 0 {
		allErrors = append(allErrors, err...)
	}

	if len(allErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": allErrors,
		})
		return
	}

	// Parse date strings to time.Time
	DateOfEvent, err := req.ParseDates()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid date format",
		})
		return
	}

	// create schedule id first

	// create event
	var event models.EventSchedule
	event.EventScheduleName = req.EventName
	event.EventScheduleDate = DateOfEvent
	// TODO: add ScheduleID from create schedule
	// event.ScheduleID = 
	err = event.CreateEventSchedule()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, event)
}
