package suprlib

import (
	"io/ioutil"
)

func ReadJsonFile(filename string) (map[string]interface{}, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("No File")
	}
	return JsonToMap(f)
}
