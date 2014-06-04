/*
Command line interface for extracting timestamps from MTS video files.
If a timestamp is not found, the filename (minus extension) is printed.

Usage:

	extract_mts_timestamp 00134.MTS ["2006-01-02 15:04:05"]
*/
package main

import (
	"fmt"
	"github.com/homburg/mtstimestamp"
	"io"
	"os"
	"path/filepath"
)

// "Return" to filename wo ext to stdout
// if a valid timestamp cannot be found
func writeFilenameWoExtAndExit(filename string) {
	extension := filepath.Ext(filename)
	filenameWoExt := filename[0 : len(filename)-len(extension)]
	fmt.Printf(string(filenameWoExt))
	os.Exit(1)
}

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
		fmt.Fprint(os.Stderr, err)
		writeFilenameWoExtAndExit(filename)
	}

	ts, err := mtstimestamp.Extract(f)

	if nil != err {
		fmt.Fprintln(os.Stderr, err)
		writeFilenameWoExtAndExit(filename)
	}

	if nil == ts {
		fmt.Fprintln(os.Stderr, "Could not get timestamp, returned original filename.")
		writeFilenameWoExtAndExit(filename)
	}

	fmt.Printf(ts.Format(format))
}
