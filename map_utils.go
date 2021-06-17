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

func CloneMapInt64Int64(data map[int64]int64) map[int64]int64 {
	m := make(map[int64]int64)
	for k, v := range data {
		m[k] = v
	}
	return m
}

func MapInt64Int64Keys(data map[int64]int64) []int64 {
	s := make([]int64, 0)
	for k := range data {
		s = append(s, k)
	}
	return s
}

func ContainsKey(m map[string]interface{}, key string) bool {
	_, ok := m[key]
	return ok
}

func ContainsKeys(m map[string]interface{}, keys []string) bool {
	for _, key := range keys {
		_, ok := m[key]
		if !ok {
			return false
		}
	}
	return true
}
