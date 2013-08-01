// TODO: Package doc
package mtstimestamp

import (
	"fmt"
	"bytes"
	"strconv"
	"time"
	"io"
)

var firstBytePattern = []byte{0x66}
var secondBytePattern = []byte{0x4D, 0x44, 0x50, 0x4D}

// Extract mts timestamp from reader
// The timestamp is assumed to be in "Local" Location/timezone
func Extract(file io.ReadSeeker) (*time.Time, error) {

	bufOne := make([]byte, 1)
	bufFour := make([]byte, 4)

	for {
		_, err := file.Read(bufOne)

		if nil != err {
			return nil, err
		}

		// Search for first single-byte pattern
		if 0 == bytes.Compare(bufOne, firstBytePattern) {
			_, err := file.Read(bufFour)

			if nil != err {
				return nil, err
			}

			// Search for 4-byte pattern
			if 0 == bytes.Compare(bufFour, secondBytePattern) {

				// Jump 3 bytes
				file.Seek(3, 1)

				// Read 8 bytes
				var dateData = make([]byte, 8)
				_, err := file.Read(dateData)

				if nil != err {
					return nil, err
				}

				// Take the hexadecimal number as date parts
				// eg. []byte{0x20, 0x13} -> "2013"
				convertedBytes := make([]string, 8)
				for i, b := range dateData {
					convertedBytes[i] = fmt.Sprintf("%x", b)
				}

				year, err := strconv.Atoi(convertedBytes[0] + convertedBytes[1])
				if nil != err {
					return nil, err
				}
				month, err := strconv.Atoi(convertedBytes[2])
				if nil != err {
					return nil, err
				}
				day, err := strconv.Atoi(convertedBytes[4])
				if nil != err {
					return nil, err
				}
				hour, err := strconv.Atoi(convertedBytes[5])
				if nil != err {
					return nil, err
				}
				minute, err := strconv.Atoi(convertedBytes[6])
				if nil != err {
					return nil, err
				}
				second, err := strconv.Atoi(convertedBytes[7])
				if nil != err {
					return nil, err
				}

				// The Location/timezone should be the local timezone
				// for the given date
				timestamp := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)

				return &timestamp, nil
			}
		}
	}
}
