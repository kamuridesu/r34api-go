package main

import (
	"fmt"
	"io"
	"net/http"
)

func sendGETRequest(url string) (string, error) {
	fmt.Printf("Request URL: %s\n", url)
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if !(res.StatusCode >= 200 && res.StatusCode <= 300) {
		return "", fmt.Errorf("response code is not in expected range! Expected %d, got %d", 200, res.StatusCode)
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}
