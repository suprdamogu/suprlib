package suprlib

import (
	"io"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	req.Close = true
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func HttpGetAsync(url string, headers map[string]string, callback func([]byte, error)) {
	go func() {
		s, err := HttpGet(url, headers)
		callback(s, err)
	}()
}

func HttpPost(url string, headers map[string]string, body io.Reader) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	req.Close = true
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func HttpPostAsync(url string, headers map[string]string, body io.Reader, callback func([]byte, error)) {
	go func() {
		s, err := HttpPost(url, headers, body)
		callback(s, err)
	}()
}

func HttpDelete(url string, headers map[string]string, body io.Reader)  ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, body)
	if err != nil {
		return nil, err
	}
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	req.Close = true
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
