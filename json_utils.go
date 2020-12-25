package suprlib

import (
	"bytes"
	"encoding/json"
)

func Z_JsonToMap(s []byte) (map[string]interface{}, error) {
	var msgMap map[string]interface{}
	decoder := json.NewDecoder(bytes.NewBuffer(s))
	decoder.UseNumber()

	err := decoder.Decode(&msgMap)
	if err != nil {
		return nil, err
	}
	return msgMap, nil
}
