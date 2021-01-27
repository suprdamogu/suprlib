package suprlib

import "strconv"

func I64Str(num int64) string {
	return strconv.FormatInt(num, 10)
}

func F64Str(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}

func StrF64(s string) float64 {
	f64, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic("Str to F64 Error,s:" + s)
	}
	return f64
}

func StrI64(s string) int64 {
	I64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic("Str to I64 Error,s:" + s)
	}
	return I64
}
