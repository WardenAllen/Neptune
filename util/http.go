package util

import (
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	postMethod = "POST"
)

func HttpPost(url string, body string) (data []byte, status int, err error) {

	var (
		payload		*strings.Reader
		client		*http.Client
		req			*http.Request
		res			*http.Response
	)

	payload = strings.NewReader(body)

	client = &http.Client{}

	req, err = http.NewRequest(postMethod, url, payload)
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/json")

	res, err = client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	data, err = ioutil.ReadAll(res.Body)
	status = res.StatusCode

	return

}



