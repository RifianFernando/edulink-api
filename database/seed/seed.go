package seed

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/edulink-api/connections"
	"github.com/go-playground/validator/v10"
)

// Validate function performs validation and insertion for a slice of models.
func Validate(models interface{}) error {
	validate := validator.New()

	// Ensure `models` is a slice using reflection.
	slice := reflect.ValueOf(models)
	if slice.Kind() != reflect.Slice {
		return errors.New("provided input is not a slice")
	}

	// Check if the slice is empty.
	if slice.Len() == 0 {
		return errors.New("no data to seed")
	}

	// Aggregate errors.
	var errs []string

	// Iterate over the items in the slice.
	for i := 0; i < slice.Len(); i++ {
		// Take a pointer to the item for database insertion.
		item := slice.Index(i).Addr().Interface()

		// Perform validation.
		if err := validate.Struct(item); err != nil {
			errs = append(errs, fmt.Sprintf("Validation failed for item %d: %v", i + 1, err))
			continue // Skip saving this item if validation fails.
		}

		// Insert the item into the database.
		if result := connections.DB.Create(item); result.Error != nil {
			errs = append(errs, fmt.Sprintf("Failed to insert item %d: %v", i, result.Error))
		}
	}

	// Return aggregated errors, if any.
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "; "))
	}

	return nil
}

// Seed function to seed all necessary data.
func Seed() error {
	fmt.Println("Starting database seeding...")

	// Aggregate errors from all seeders.
	var errs []string

	if err := Validate(UserSeeder()); err != nil {
		errs = append(errs, fmt.Sprintf("UserSeeder error: %v", err))
	}
	if err := Validate(TeacherSeeder()); err != nil {
		errs = append(errs, fmt.Sprintf("TeacherSeeder error: %v", err))
	}
	if err := Validate(GradeSeeder()); err != nil {
		errs = append(errs, fmt.Sprintf("GradeSeeder error: %v", err))
	}
	if err := Validate(ClassSeeder()); err != nil {
		errs = append(errs, fmt.Sprintf("ClassSeeder error: %v", err))
	}
	if err := Validate(StudentSeeder()); err != nil {
		errs = append(errs, fmt.Sprintf("StudentSeeder error: %v", err))
	}
	if err := Validate(AdminSeeder()); err != nil {
		errs = append(errs, fmt.Sprintf("AdminSeeder error: %v", err))
	}
	if err := Validate(StaffSeeder()); err != nil {
		errs = append(errs, fmt.Sprintf("StaffSeeder error: %v", err))
	}
	if err := Validate(SubjectSeeder()); err != nil {
		errs = append(errs, fmt.Sprintf("SubjectSeeder error: %v", err))
	}
	if err := Validate(AttendanceSeeder()); err != nil {
		errs = append(errs, fmt.Sprintf("SubjectSeeder error: %v", err))
	}

	// Return aggregated errors, if any.
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "; "))
	}

	fmt.Println("Database seeding completed!")
	return nil
}
