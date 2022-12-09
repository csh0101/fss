package utils

import "strconv"

func Str2Uint64(input string) (uint64, error) {
	return strconv.ParseUint(input, 10, 64)
}
