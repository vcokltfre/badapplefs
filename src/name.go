package src

import (
	"strconv"
)

var (
	space     = ' '
	zeroSpace = "\u200b"
)

func CreateName(name int) string {
	// binary := strconv.FormatInt(int64(name), 2)

	return strconv.Itoa(name)
	// if os.Getenv("NAME_SIMPLE") == "1" {
	// }

	// var result string

	// for i := 0; i < len(binary); i += 1 {
	// 	if binary[i] == '1' {
	// 		result += "'"
	// 	} else {
	// 		result += ","
	// 	}
	// }

	// return result
}
