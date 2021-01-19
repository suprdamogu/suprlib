package suprlib

import "time"

func TimeStr(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
