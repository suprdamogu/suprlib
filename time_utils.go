package suprlib

import "time"

func TimeStr(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func TimeStamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

