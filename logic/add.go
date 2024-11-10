package logic

import (
	"bufio"
	"cli-to-do-list/constants"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func AddRecord() {
	// Get last record ID from task list
	lastRecord, err := findPreviousTaskID()

	// Convert string to int so ID can be increased by 1
	lastRecordInt, err := strconv.Atoi(lastRecord)
	if err != nil {
		log.Fatal(err)
	}

	// Increase ID by 1
	lastRecordInt++

	// Convert int back to string
	lastRecord = strconv.Itoa(lastRecordInt)

	// Open file to add a new task
	f, err := os.OpenFile(constants.GetTaskDataFile(), os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}

	// Schedule file close at the end of the function
	defer f.Close()

	// Assign a writer to write the file
	w := csv.NewWriter(f)

	// Assign Scanner to get input data from user
	scanner := bufio.NewScanner(os.Stdin)

	// Retrieve task description from user input
	fmt.Println("Enter the to-do description:")
	scanner.Scan()
	description := scanner.Text()
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve task due date from user input
	fmt.Println("Enter the to-do due-date:")
	scanner.Scan()
	due_date := scanner.Text()
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	// Create a record containing task data
	record := []string{
		lastRecord, description, due_date,
	}
	if err := w.Write(record); err != nil {
		log.Fatalf("Error writing record: %v", err)
	}

	// Flush buffered data to file
	w.Flush()

	// Check for any errors during the flush
	if err := w.Error(); err != nil {
		log.Fatalf("Error flushing writer: %v", err)
	}
}

// Function to find the ID of the last task in the list
func findPreviousTaskID() (string, error) {
	f, err := os.Open(constants.GetTaskDataFile())
	if err != nil {
		log.Fatal(err)
	}

	// Assign a reader to read from the file
	r := csv.NewReader(f)
	var lastRecord string

	// Iterate through task records
	for {
		rec, err := r.Read()

		// If end of file is reached return the last record ID
		if err == io.EOF {
			return lastRecord, nil
		}
		if err != nil {
			return "", nil
		}
		// Assign current row ID as the last record in case next row
		// is the end of the file
		lastRecord = rec[0]
	}
}
