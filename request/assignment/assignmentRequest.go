package request

import (
	"github.com/edulink-api/database/models"
	"github.com/edulink-api/res"
	req "github.com/edulink-api/request"
)

/*
* InsertAssignmentRequest struct
* see the documentation about binding and validation here:
* https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/#validating-date-time
 */
type InsertAssignmentRequest struct {
	models.Assignment
}

// Validate method
func (r *InsertAssignmentRequest) Validate() []map[string]string {

	// Validate the struct
	err := res.ResponseMessage(req.Validate.Struct(r))

	return err
}
