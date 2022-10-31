package utils

import (
	"io"
	"net/http"
	"time"
)

func ApiRequest(url string) (io.ReadCloser, error) {
	client := &http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "go-test")

	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 200 {
		return nil, err
	}

	return resp.Body, err
}
