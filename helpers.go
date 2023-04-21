package utopiapaylib

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func closeRequest(resp *http.Response) {
	if resp != nil {
		resp.Body.Close()
	}
}

func sendRequest(requestType, url string, requestData []byte) ([]byte, error) {
	req, err := http.NewRequest(requestType, url, bytes.NewBuffer(requestData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	defer closeRequest(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}
