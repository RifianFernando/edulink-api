package helper

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/edulink-api/database/models"
)

var (
	AcademicStartSemesterMonth = 7 // July
)

func GetOrCreateAcademicYear() (models.AcademicYear, error) {
	// get year now
	var academicSemesterYear string
	yearNow := time.Now().Year()

	// check if the semester 1 or 2 from the month
	if int(time.Now().Month()) >= AcademicStartSemesterMonth {
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

func GetAcademicYearList() ([]models.AcademicYear, error) {
	// get all academic year
	academicYear := models.AcademicYear{}

	academicYearList, err := academicYear.GetAcademicYearList()
	if err != nil {
		return academicYearList, err
	}

	return academicYearList, nil
}

func ValidateAcademicYearInput(academicYear string) error {
	if academicYear == "" {
		return fmt.Errorf("academic_year is required")
	}

	academicYearStart := strings.Split(academicYear, "/")[0]
	academicYearEnd := strings.Split(academicYear, "/")[1]

	parsedIntAcademicYearStart, err := strconv.ParseInt(academicYearStart, 10, 64)
	if err != nil {
		return fmt.Errorf("academic_year_start must be a number")
	}

	parsedIntAcademicYearEnd, err := strconv.ParseInt(academicYearEnd, 10, 64)
	if err != nil {
		return fmt.Errorf("academic_year_end must be a number")
	}
	fmt.Println("rizz:", academicYearStart, academicYearEnd)
	if academicYearStart == "" || academicYearEnd == "" {
		return fmt.Errorf("academic_year_start and academic_year_end are required")
	} else if academicYearStart == academicYearEnd {
		return fmt.Errorf("academic_year_start and academic_year_end cannot be the same")
	} else if academicYearStart > academicYearEnd {
		return fmt.Errorf("academic_year_start cannot be greater than academic_year_end")
	} else if (parsedIntAcademicYearStart + 1) != parsedIntAcademicYearEnd {
		return fmt.Errorf("academic_year_end must be exactly 1 year greater than academic_year_start")
	} else if int(time.Now().Month()) < AcademicStartSemesterMonth {
		if int(parsedIntAcademicYearStart) >= time.Now().Year() {
			return fmt.Errorf("academic_year_start must be exactly 1 year less than the current year because the new semester starts in July")
		}
	}

	return nil
}
