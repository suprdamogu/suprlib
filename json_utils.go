package suprlib

import (
	"bytes"
	"encoding/json"
)

func JsonMarshal(v interface{}) []byte {
	j, err := json.Marshal(v)
	if err != nil {
		panic(DEAD_CODE)
	}
	return j
}

func JsonToMap(s []byte) (map[string]interface{}, error) {
	var msgMap map[string]interface{}
	decoder := json.NewDecoder(bytes.NewBuffer(s))
	decoder.UseNumber()

	err := decoder.Decode(&msgMap)
	if err != nil {
		return nil, err
	}
	return msgMap, nil
}

func JsonToObj(s []byte, out interface{}) error {
	decoder := json.NewDecoder(bytes.NewBuffer(s))
	decoder.UseNumber()
	err := decoder.Decode(out)
	if err != nil {
		return err
	}
	return nil
}
