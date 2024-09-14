package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func GetEnvValue(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s is not set", key)
	} else if value == "" {
		log.Fatalf("Environment variable %s is empty", key)
	}
	return value
}

func SetEnvValue(key, value string) (string, error) {
	// Load the existing .env file
	file, err := os.Open(".env")
	if err != nil {
		return "", fmt.Errorf("error opening .env file: %v", err)
	}
	defer file.Close()

	// Read the file line by line
	var lines []string
	scanner := bufio.NewScanner(file)
	found := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, key+"=") {
			// Replace the existing key-value pair
			lines = append(lines, key+"="+value)
			found = true
		} else {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading .env file: %v", err)
	}

	// If the key was not found, append it to the end of the file
	if !found {
		lines = append(lines, key+"="+value)
	}

	// Write the modified contents back to the .env file
	err = os.WriteFile(".env", []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		return "", fmt.Errorf("error writing to .env file: %v", err)
	}

	return value, nil
}
