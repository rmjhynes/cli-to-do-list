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

func AddRecord() {
	findPreviousTaskID()

	f, err := os.OpenFile(constants.TaskData, os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}

	// Schedule file close at the end of the function
	defer f.Close()

	w := csv.NewWriter(f)

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

	record := []string{
		"1", description, due_date,
	}
	if err := w.Write(record); err != nil {
		log.Fatalf("Error writing record: %v", err)
	}

	w.Flush()

	// Check for any errors during the flush
	if err := w.Error(); err != nil {
		log.Fatalf("Error flushing writer: %v", err)
	}

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

func findPreviousTaskID() (string, error) {
	f, err := os.Open(constants.TaskData)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)
	var lastRecord string

	for {
		rec, err := r.Read()

		if err == io.EOF {
			fmt.Println("The last record: ", lastRecord)
			return lastRecord, nil
		}
		if err != nil {
			return "", nil
		}
		lastRecord = rec[0]
	}
}
