package suprlib

import (
	"fmt"
	"testing"
)

func TestZzzCrypto_MD5(t *testing.T) {
	md5 := ZCrypto.MD5([]byte("123456789"))
	fmt.Println(md5)
}
