package migration

import (
	"time"

	"github.com/edulink-api/database/migration/lib"
)

/*
* see the documentation here
* https://gorm.io/docs/data_types.html
* https://gorm.io/docs/models.html#Fields-Tags
 */
type AttendanceStatus string

const (
	Present AttendanceStatus = "Present"
	Sick    AttendanceStatus = "Sick"
	Leave   AttendanceStatus = "Leave"
	Absent  AttendanceStatus = "Absent"
)

type Attendance struct {
	AttendanceID     int64            `gorm:"primaryKey;autoIncrement"`
	StudentID        int64            `gorm:"not null"`
	ClassNameID      int64            `gorm:"not null"`
	AttendanceDate   time.Time        `gorm:"not null"`
	AttendanceStatus AttendanceStatus `gorm:"not null;type:VARCHAR(8)"`
	lib.BaseModel                     /* this type include CreatedAt, UpdatedAt, DeletedAt, I can't use the gorm.models because can't customize the id name */
}

/*
* see the documentation here about conventions
* https://gorm.io/docs/conventions.html
 */
func (Attendance) TableName() string {
	return lib.GenerateTableName(lib.Administration, "attendances")
}
