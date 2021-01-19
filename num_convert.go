package suprlib

import "strconv"

func I64Str(num int64) string {
	return strconv.FormatInt(num, 10)
}

func F64Str(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}
