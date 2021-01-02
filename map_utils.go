package suprlib

import "encoding/json"

func JsonMapGetInt64Value(data map[string]interface{}, key string) (int64, error) {
	vT, ok := data[key]
	if !ok {
		return 0, MAP_KEY_NOT_FOUND_ERROR
	}
	vNum, ok := vT.(json.Number)
	if !ok {
		return 0, JSON_MAP_TRANS_ERROR
	}
	return vNum.Int64()
}

func JsonMapGetUIntValue(data map[string]interface{}, key string) (uint, error) {
	vI64, err := JsonMapGetInt64Value(data, key)
	if err != nil {
		return 0, err
	}
	return uint(vI64), nil
}
