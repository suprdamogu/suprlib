package suprlib

import (
	"io/ioutil"
	"os"
)

func FileExists(filename string) bool{
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func ReadFile(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("No File")
	}
	return b
}

func ReadJsonMapFile(filename string) (map[string]interface{}, error) {
	b := ReadFile(filename)
	return JsonToMap(b)
}

func ReadJsonArrayFile(filename string) (interface{}, error) {
	b := ReadFile(filename)
	out := make([]interface{}, 0)
	err := JsonToObj(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

