package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// CallRemoteEndpoint calls a remote endpoint and returns the response
// body as a string
func CallRemoteEndpoint(URL string) (string, error) {
	res, err := http.Get(URL)
	if err != nil {
		return "", fmt.Errorf("Couldn't make http request: %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		return "", fmt.Errorf("Remote endpoint returned status %d", res.StatusCode)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("Failed to read response body: %v", err)
	}
	return string(body), nil
}
