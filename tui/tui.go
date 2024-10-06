package tui

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// get todo list data from csv file
// print to the console in tabular format

func ListRecords() {
	f, err := os.Open("/Users/rmjhynes/devops/golang/cli-to-do-list/tui/data.csv")
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
		fmt.Printf("%+v\n", rec)
	}
}
