package http

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func PostJSON(url string, body string) (string, error) {
	var (
		res *http.Response
		err error
	)
	res, err = http.Post(url, "application/json;charset=utf-8", strings.NewReader(body))
	if err != nil || res == nil {
		return "", err
	}
	defer func() {
		if res != nil && res.Body != nil {
			res.Body.Close()
		}
	}()
	var content []byte
	content, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
