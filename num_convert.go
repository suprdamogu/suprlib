package suprlib

import "strconv"

func I64ToStr(num int64) string {
	return strconv.FormatInt(num, 10)
}

func F64ToStr(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}
