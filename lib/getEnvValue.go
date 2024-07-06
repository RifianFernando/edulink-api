package lib

import (
	"log"
	"os"
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
