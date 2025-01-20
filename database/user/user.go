package user

import "github.com/gin-gonic/gin"

var (
	Admin string = "admin"
	Staff string = "staff"
	Teacher string = "teacher"
	HomeRoomTeacher string = "homeroom_teacher"
)

func CheckUserRole(userType []string, role string) bool {
	for _, r := range userType {
		if r == role {
			return true
		}
	}

	return false
}

func ValidateUserRoleCtx(c *gin.Context, role string) bool {
	userRole, exist := c.Get("user_type")
	if !exist {
		return false
	}

	return CheckUserRole(userRole.([]string), role)
}

func GetUserTypeFromCtx(c *gin.Context) []string {
	userType, exist := c.Get("user_type")
	if !exist {
		return nil
	}

	return userType.([]string)
}
