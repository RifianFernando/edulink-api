package controllers

import (
	"net/http"
	"strconv"

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

	DateEventStartParsed := req.EventDateStart
	DateEventEndParsed := req.EventDateEnd

	// create event
	var event models.EventSchedule
	event.EventScheduleName = req.EventName
	event.EventScheduleDateStart = DateEventStartParsed
	event.EventScheduleDateEnd = DateEventEndParsed

	err := event.CreateEventSchedule()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, event)
}

func UpdateEvent(c *gin.Context) {
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

	// get event id from url
	eventID := c.Param("event_id")

	DateEventStartParsed := req.EventDateStart
	DateEventEndParsed := req.EventDateEnd

	// create event
	var event models.EventSchedule
	var err error
	event.EventScheduleID, err = strconv.ParseInt(eventID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.EventScheduleName = req.EventName
	event.EventScheduleDateStart = DateEventStartParsed
	event.EventScheduleDateEnd = DateEventEndParsed

	err = event.UpdateEventSchedule()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, event)
}

func DeleteEvent(c *gin.Context) {
	// get event id from url
	eventID := c.Param("event_id")

	// create event
	var event models.EventSchedule
	var err error
	event.EventScheduleID, err = strconv.ParseInt(eventID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = event.DeleteEventSchedule()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "Event deleted successfully",
	})
}
