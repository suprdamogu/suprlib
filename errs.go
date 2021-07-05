package suprlib

import "errors"

var MAP_KEY_NOT_FOUND_ERROR = errors.New("MAP_KEY_NOT_FOUND_ERROR")
var JSON_MAP_TRANS_ERROR = errors.New("JSON_MAP_TRANS_ERROR")
var NOT_FOUND_ERROR = errors.New("NOT_FOUND_ERROR")
var DEAD_CODE = errors.New("DEAD_CODE")

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}
