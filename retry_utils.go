package suprlib

import "time"

// Retry TODO 类型不好控制，以后支持泛型再说吧
func Retry(f func() error, counts int, delay time.Duration) error {
	var err error
	for i := 0; i < counts; i++ {
		err = f()
		if err != nil {
			time.Sleep(delay)
			continue
		}
		break
	}
	return err
}
