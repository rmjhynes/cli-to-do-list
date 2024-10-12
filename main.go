package main

import (
	"cli-to-do-list/cmd"
	"fmt"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
