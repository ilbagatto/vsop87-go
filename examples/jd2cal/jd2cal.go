// Converts Julian date to civil date.
// Usage:
//
//	jd2cal JD
//
// JD is a julian date, number of days elapsed since mean UT noon
// of January 1st 4713 BC. e.g. 2460047.86458333
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ilbagatto/vsop87-go/internal/timeutils"
)

func main() {
	var jd float64
	var err error

	if len(os.Args) > 1 {
		jdStr := os.Args[1]
		jd, err = strconv.ParseFloat(jdStr, 32)
		if err != nil {
			fmt.Printf("Invalid Julian Date: %s\n.", jdStr)
			os.Exit(1)
		}

	} else {
		fmt.Print("Usage: jd2cal JD\n")
		os.Exit(1)
	}

	dateStr := timeutils.JulianToDateString(jd)
	fmt.Println(dateStr)
	os.Exit(0)
}
