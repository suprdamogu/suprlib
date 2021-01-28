package suprlib

import "time"

func TimeStr(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func TimeToTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func StrTime(s string) time.Time {
	dt, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		panic("Str to Time Error,s:" + s)
	}
	return dt
}

func TimeFromTimestamp(timestamp int64) time.Time {
	return time.Unix(timestamp/1000, (timestamp%1000)*1e6)
}
