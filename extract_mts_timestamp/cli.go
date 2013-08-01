package main

import (
	"os"
	"fmt"
	"github.com/homburg/mtstimestamp"
)

func main() {
	filename := ""
	if len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
		fmt.Printf("Usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}

	format := "2006-01-02 15.04.05"
	if len(os.Args) > 2 {
		format = os.Args[2]
	}

	f, err := os.Open(filename)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	ts, err := mtstimestamp.Extract(f)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	if nil == ts {
		fmt.Println("Could not get timestamp")
		os.Exit(1)
	}

	fmt.Printf(ts.Format(format))
}
