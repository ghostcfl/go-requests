package requests

import (
	"encoding/json"
	"io"
	"net/url"
	"strings"
)

func MapToJsonString(m J) (string, error) {
	jsonString, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(jsonString), nil
}

func MapToQueryString(m KV) string {
	// Create a new url.Values object
	values := url.Values{}

	// Loop through the map and add each key-value pair to the url.Values object
	for key, value := range m {
		values.Add(key, value)
	}

	// Encode the url.Values object as a query string and return it
	return strings.Replace(values.Encode(), "+", "%20", -1)
}

func stringToReadCloser(s string) io.ReadCloser {
	rc := io.NopCloser(strings.NewReader(s))
	return rc
}
