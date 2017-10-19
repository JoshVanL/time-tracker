package main

import (
	"fmt"

	"github.com/joshvanl/time-tracker/pkg/timesheet"
)

func main() {

	t, err := timesheet.New()
	if err != nil {
	}

	//if err := t.WriteConfig(); err != nil {
	//	fmt.Printf("ERROR: %v", err)
	//}

	err = t.ReadConfig()
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}

	//fmt.Printf("%v", t)
	//fmt.Printf("%v", t.Stamps()[0])
	//fmt.Printf("%s", t.FilePath())

	// Read form a file
	// Get time from some time period
	// Apply time tracking
	// Comments?
	// Ammend
}
