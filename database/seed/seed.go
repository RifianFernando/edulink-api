package seed

import (
	"fmt"
	"reflect"

	"github.com/edulink-api/connections"
	"github.com/go-playground/validator/v10"
)

// Validate function takes a slice of any model and performs validation and insertion.
func Validate(models interface{}) {
	validate := validator.New()

	// Use reflection to ensure `models` is a slice
	slice := reflect.ValueOf(models)
	if slice.Kind() != reflect.Slice {
		fmt.Println("Validate function requires a slice")
		return
	}

	// Check if the slice is empty
	if slice.Len() == 0 {
		fmt.Println("Empty slice provided, nothing to validate")
		return
	}

	// Iterate over each item in the slice
	for i := 0; i < slice.Len(); i++ {
		// Assert that each item is a struct and take a pointer to it
		item := slice.Index(i).Addr().Interface()

		// TODO: fix the validation gender if lowercase
		if err := validate.Struct(item); err != nil {
			fmt.Printf("Validation failed for item %d: %v\n", i, err)
			continue // Skip saving this item if validation fails
		}

		// Insert into the database if validation passes
		if result := connections.DB.Create(item); result.Error != nil {
			fmt.Printf("Failed to insert item %d: %v\n", i, result.Error)
		} else {
			fmt.Printf("Item %d inserted successfully!\n", i)
		}
	}
}

// Seed function to call each seeder with validation
func Seed() {
	Validate(UserSeeder())
	Validate(TeacherSeeder())
	Validate(GradeSeeder())
	Validate(ClassSeeder())
	Validate(StudentSeeder())
	Validate(AdminSeeder())
	Validate(SubjectSeeder())
}
