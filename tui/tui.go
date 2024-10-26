package tui

import (
	"cli-to-do-list/constants"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/tabwriter"
)

// get todo list data from csv file
// print to the console in tabular format

func ListRecords() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)

	f, err := os.Open(constants.TaskData)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)

	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// Join the record (slice of strings) into a single string with tab separators
		line := strings.Join(rec, "\t")

		// Write the tab-separated line to the tab writer
		fmt.Fprintln(w, line)
	}

	// Flush the writer to ensure all output is written
	w.Flush()
}
