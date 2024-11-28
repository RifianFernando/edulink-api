package models

import (
	"fmt"
	"time"

	"github.com/edulink-api/connections"
	"github.com/edulink-api/database/migration/lib"
	"gorm.io/gorm"
)

type Attendance struct {
	AttendanceID     int64     `gorm:"primaryKey;autoIncrement"`
	StudentID        int64     `gorm:"not null"`
	AttendanceDate   time.Time `gorm:"not null"`
	AttendanceStatus string    `gorm:"not null" validate:"oneof='Absent' 'Leave' 'Sick' 'Present'"`
}

type AttendanceModel struct {
	Attendance
	Student Student `gorm:"foreignKey:StudentID;references:StudentID"`
}

func (Attendance) TableName() string {
	return lib.GenerateTableName(lib.Administration, "attendances")
}

func GetAllAttendanceMonthSummaryByClassID(classID string, date time.Time) (interface{}, error) {
	type AttendanceStats struct {
		Date         time.Time `json:"date"`
		PresentTotal int       `json:"present_total"`
		SickTotal    int       `json:"sick_total"`
		LeaveTotal   int       `json:"leave_total"`
		AbsentTotal  int       `json:"absent_total"`
	}
	var attendanceStats []AttendanceStats
	err := connections.DB.Model(Attendance{}).
		Select("attendance_date AS date, "+
			"SUM(CASE WHEN attendance_status = 'Present' THEN 1 ELSE 0 END) AS present_total, "+
			"SUM(CASE WHEN attendance_status = 'Sick' THEN 1 ELSE 0 END) AS sick_total, "+
			"SUM(CASE WHEN attendance_status = 'Leave' THEN 1 ELSE 0 END) AS leave_total, "+
			"SUM(CASE WHEN attendance_status = 'Absent' THEN 1 ELSE 0 END) AS absent_total").
		Joins("JOIN academic.students s ON attendances.student_id = s.student_id").
		Where("EXTRACT(YEAR FROM attendance_date) = ? AND EXTRACT(MONTH FROM attendance_date) = ? AND s.class_name_id = ?", date.Year(), int(date.Month()), classID).
		Group("attendance_date").
		Order("attendance_date DESC").
		Scan(&attendanceStats).Error

	if err != nil {
		return nil, err
	}

	return attendanceStats, nil
}

func GetAllStudentAttendanceDateByClassID(classID string, date time.Time) (interface{}, error) {
	targetDate := date.Truncate(24 * time.Hour)
	type AttendanceStats struct {
		ID     int64     `json:"id"`
		Name   string    `json:"name"`
		Sex    string    `json:"sex"`
		Reason string    `json:"reason"`
		Date   time.Time `json:"date"`
	}
	var attendanceStats []AttendanceStats
	err := connections.DB.Model(Attendance{}).
		Select(
			"s.student_id AS id, "+
				"s.student_name AS name, "+
				"s.student_gender AS sex, "+
				"attendances.attendance_status AS reason, "+
				"attendances.attendance_date AS date",
		).
		Joins("JOIN academic.students s ON s.student_id = attendances.student_id").
		Where(
			"s.class_name_id = ? AND "+
			"EXTRACT(YEAR FROM attendances.attendance_date) = ? AND "+
				"EXTRACT(MONTH FROM attendances.attendance_date) = ? AND "+
				"EXTRACT(DAY FROM attendances.attendance_date) = ?", classID,
			targetDate.Year(), int(targetDate.Month()), targetDate.Day(),
		).
		Order("s.student_name ASC").
		Find(&attendanceStats).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return AttendanceModel{}, fmt.Errorf("no students found")
		}
		return AttendanceModel{}, err
	}

	return attendanceStats, nil
}
