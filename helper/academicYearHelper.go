package helper

import (
	"strconv"
	"time"

	"github.com/edulink-api/database/models"
)

func GetOrCreateAcademicYear() (models.AcademicYear, error) {
	// get year now
	var academicSemesterYear string
	yearNow := time.Now().Year()

	// check if the semester 1 or 2 from the month
	if time.Now().Month() >= 7 {
		// concatenate the academic year
		academicSemesterYear = strconv.Itoa(yearNow) + "/" + strconv.Itoa(yearNow+1)
	} else {
		academicSemesterYear = strconv.Itoa(yearNow-1) + "/" + strconv.Itoa(yearNow)
	}

	// search for the assignment academic year if not exist create it
	var academicYear models.AcademicYear
	academicYear.AcademicYear = academicSemesterYear

	err := academicYear.GetAcademicYearByModel()
	if err != nil || academicYear.AcademicYearID == 0 {
		err = academicYear.CreateAcademicYear()
		if err != nil {
			return academicYear, err
		}
	}

	return academicYear, nil
}
