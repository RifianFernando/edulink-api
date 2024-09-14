package lib

// bcrypt password hashing
import (
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func CompareHash(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return nil
	}

	return err
}

// GenerateRandomString generates a random string of the given length from the charset.
func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

// GenerateMultipleRandomStrings generates `n` random strings, each of the specified length.
func GenerateMultipleRandomStrings(n, length int) []string {
	randomStrings := make([]string, n)
	for i := 0; i < n; i++ {
		randomStrings[i] = GenerateRandomString(length)
	}
	return randomStrings
}
