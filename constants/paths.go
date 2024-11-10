package constants

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Get the path to the data file from .env file
func GetTaskDataFile() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	TaskData := os.Getenv("TASK_DATA_FILE")

	return TaskData
}
