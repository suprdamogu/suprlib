package suprlib

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

func FileExists(filename string) bool {
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

func WriteJsonFile(filename string, obj interface{}) error {
	b, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, b, 0666)
	return err
}

func AppendToFile(filename string, b []byte) error {
	f, err := os.OpenFile(filename, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	n, err := f.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}
	_, err = f.WriteAt(b, n)
	defer f.Close()
	return nil
}
