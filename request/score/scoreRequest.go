package request

import (
	"fmt"
	"strconv"
	"strings"

	req "github.com/edulink-api/request"
	"github.com/edulink-api/res"
)

/*
* InsertStudentScoreRequest struct
* see the documentation about binding and validation here:
* https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/#validating-date-time
 */
type InsertStudentScoreRequest struct {
	StudentID int64 `json:"studentID" binding:"required"`
	Score     int   `json:"score" binding:"required"`
}

// Validate method
func (r *InsertStudentScoreRequest) Validate() []map[string]string {

	// Validate the struct
	err := res.ResponseMessage(req.Validate.Struct(r))

	return err
}

/*
* InsertAllStudentScoreRequest struct
* see the documentation about binding and validation here:
* https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples
 */
type InsertAllStudentScoreRequest struct {
	AssignmentID              int64                       `json:"assignment_id" binding:"required"`
	InsertStudentScoreRequest []InsertStudentScoreRequest `json:"scores" binding:"required"`
}

// Validate method
func (r *InsertAllStudentScoreRequest) ValidateAllStudentScore() []map[string]string {
	// Validate the struct
	var allErrors []map[string]string
	err := res.ResponseMessage(req.Validate.Struct(r))
	if err != nil {
		return err
	}

	for i, data := range r.InsertStudentScoreRequest {
		if err := data.Validate(); err != nil {
			// index with error
			errorMap := map[string]string{
				"row-error": fmt.Sprintf("%d", i+1),
				"field":     err[0]["field"],
				"message":   err[0]["message"],
			}
			allErrors = append(allErrors, errorMap)
		}
	}

	return allErrors
}

/*
* updateStudentScore struct
* see the documentation about binding and validation here:
* https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples
 */

type UpdateStudentScoreRequest struct {
	StudentName string `json:"StudentName" binding:"required" validate:"required,min=1"`
	SubjectName string `json:"SubjectName" binding:"required" validate:"required,min=1"`
	Scores      []struct {
		AssignmentID int64 `json:"AssignmentID" binding:"required"`
		Score        int   `json:"Score" binding:"required" validate:"required,min=0,max=100"`
	} `json:"Scores" binding:"required" validate:"required,min=1,dive"`
}

// Validate method for UpdateStudentScoreRequest
func (r *UpdateStudentScoreRequest) Validate() []map[string]string {
	var allErrors []map[string]string
	err := res.ResponseMessage(req.Validate.Struct(r))
	if err != nil {
		rowErrors, _ := strconv.Atoi(strings.Split(strings.Split(err[0]["message"], "[")[1], "]")[0])
		parsedRowErrors := rowErrors + 1
		errorMap := map[string]string{
			"row-error": strconv.Itoa(parsedRowErrors),
			"field":     err[0]["field"],
			"message":   err[0]["message"],
		}
		allErrors = append(allErrors, errorMap)
	}

	return allErrors
}
