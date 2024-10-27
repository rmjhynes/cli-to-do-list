package tui

import (
	"bufio"
	"cli-to-do-list/constants"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func RemoveRecord() {
	f, err := os.OpenFile(constants.TaskData, os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}

	// Schedule file close at the end of the function
	defer f.Close()

	scanner := bufio.NewScanner(os.Stdin)

	// Get task ID to delete
	fmt.Println("Enter the ID of the to-do to mark as complete:")
	scanner.Scan()
	targetID := scanner.Text()
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	rowToDelete, err := findRowByID(targetID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Marking row:", rowToDelete, "as complete...")

	deleteRowByID(targetID)
}

// Function to find a row by the unique ID
func findRowByID(targetID string) ([]string, error) {
	f, err := os.Open(constants.TaskData)
	if err != nil {
		log.Fatal(err)
	}
	// Create a new CSV reader
	reader := csv.NewReader(f)
	// Read the CSV file line by line
	for {
		row, err := reader.Read()
		if err != nil {
			// Stop at EOF or error
			if err.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("error reading file: %v", err)
		}

		// Check if the first column matches the target ID
		if row[0] == targetID {
			return row, nil
		}
	}

	return nil, fmt.Errorf("ID %s not found", targetID)
}

// Function to delete task using task ID
func deleteRowByID(targetID string) error {
	f, err := os.Open(constants.TaskData)
	// Create a temporary file in the same directory to store updated rows
	tempFile, err := os.Create("tempFile" + ".tmp")
	if err != nil {
		return fmt.Errorf("could not create temp file: %v", err)
	}
	defer tempFile.Close()

	// Set up CSV readers and writers
	r := csv.NewReader(f)
	w := csv.NewWriter(tempFile)

	// Stream rows from the original file to the temporary file
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error reading row: %v", err)
		}

		// Write rows that do not match the target ID
		if row[0] != targetID {
			if err := w.Write(row); err != nil {
				return fmt.Errorf("error writing row: %v", err)
			}
		}
	}

	// Flush the writer to ensure all data is written
	w.Flush()
	if err := w.Error(); err != nil {
		return fmt.Errorf("error flushing writer: %v", err)
	}

	// Close the temp file before renaming
	tempFile.Close()

	// Replace the original file with the temporary file
	if err := os.Rename(tempFile.Name(), constants.TaskData); err != nil {
		return fmt.Errorf("could not replace original file: %v", err)
	}

	return nil
}
