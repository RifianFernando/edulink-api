package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/edulink-api/connections"
	"github.com/edulink-api/database/migration/lib"
	"gorm.io/gorm"
)

type Attendance struct {
	AttendanceID     int64     `gorm:"primaryKey;autoIncrement"`
	StudentID        int64     `gorm:"not null"`
	ClassNameID      int64     `gorm:"not null"`
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
		Where("EXTRACT(YEAR FROM attendance_date) = ? AND EXTRACT(MONTH FROM attendance_date) = ? AND attendances.class_name_id = ?", date.Year(), int(date.Month()), classID).
		Group("attendance_date").
		Order("attendance_date DESC").
		Scan(&attendanceStats).Error

	if err != nil {
		return nil, err
	}

	return attendanceStats, nil
}

type AttendanceDateByClassID struct {
	ID     int64     `json:"id"`
	Name   string    `json:"name"`
	Sex    string    `json:"sex"`
	Reason string    `json:"reason"`
	Date   time.Time `json:"date"`
}

func GetAllStudentAttendanceDateByClassID(classID string, date time.Time) ([]AttendanceDateByClassID, error) {
	targetDate := date.Truncate(24 * time.Hour)
	var attendanceStats []AttendanceDateByClassID
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
			"attendances.class_name_id = ? AND "+
				"EXTRACT(YEAR FROM attendances.attendance_date) = ? AND "+
				"EXTRACT(MONTH FROM attendances.attendance_date) = ? AND "+
				"EXTRACT(DAY FROM attendances.attendance_date) = ?", classID,
			targetDate.Year(), int(targetDate.Month()), targetDate.Day(),
		).
		Order("s.student_name ASC").
		Find(&attendanceStats).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []AttendanceDateByClassID{}, fmt.Errorf("attendance not found")
		}
		return []AttendanceDateByClassID{}, err
	}

	return attendanceStats, nil
}

type ClassDateAttendanceStudent struct {
	StudentID string `json:"student_id" binding:"required" validate:"required"`
	Reason    string `json:"reason" binding:"required" validate:"required,oneof='Present' 'Sick' 'Leave' 'Absent'"`
}

func UpdateStudentAttendanceByClassIDAndDate(classID string, date time.Time, studentData []ClassDateAttendanceStudent) error {
	tx := connections.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for _, data := range studentData {
		studentID, err := strconv.ParseInt(data.StudentID, 10, 64)
		if err != nil {
			tx.Rollback()
			return err
		}

		query := `
				UPDATE administration.attendances a
				SET attendance_status = ?
				FROM academic.students s
				WHERE s.student_id = a.student_id
				AND a.class_name_id = ?
				AND a.student_id = ?
				AND EXTRACT(YEAR FROM a.attendance_date) = ?
				AND EXTRACT(MONTH FROM a.attendance_date) = ?
				AND EXTRACT(DAY FROM a.attendance_date) = ?
			`

		result := tx.Exec(query,
			data.Reason,
			classID,
			studentID,
			date.Year(),
			int(date.Month()),
			date.Day(),
		)

		if result.Error != nil || result.RowsAffected == 0 {
			tx.Rollback()
			return result.Error
		}
	}

	return tx.Commit().Error
}

func CreateStudentClassAttendance(classID string, date time.Time, studentData []ClassDateAttendanceStudent) error {
	tx := connections.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for _, data := range studentData {
		studentID, err := strconv.ParseInt(data.StudentID, 10, 64)
		if err != nil {
			tx.Rollback()
			return err
		}
		classID, err := strconv.ParseInt(classID, 10, 64)
		if err != nil {
			tx.Rollback()
			return err
		}

		attendance := Attendance{
			StudentID:        studentID,
			ClassNameID:      classID,
			AttendanceDate:   date,
			AttendanceStatus: data.Reason,
		}

		result := tx.Create(&attendance)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}

	return tx.Commit().Error
}
