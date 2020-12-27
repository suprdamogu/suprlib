package suprlib

import "time"

type (
	ZzzCrypto    byte
	ZzzHttp      byte
	ZzzJson      byte
	ZzzWebsocket byte
)

const (
	// Version 版本号
	Version = "0.0.1"
)

var (
	// ZUpTime 当前服务启动时间
	ZUpTime = time.Now()

	// ZCrypto 加密解密
	ZCrypto ZzzCrypto

	// ZHttp Http/Https相关
	ZHttp ZzzHttp

	// ZJson Json相关
	ZJson ZzzJson

	// ZWebsocket Websocket相关
	ZWebsocket ZzzWebsocket
)
