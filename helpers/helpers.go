package helpers

import "net/http"

// NilToEmptyMap is a helper function to convert nil value to {}
func NilToEmptyMap(d *interface{}) interface{} {
	data := *d
	if *d == nil {
		data = make(map[string]interface{})
	}
	return data
}

// Fetch is a helper function to get the http response
func Fetch(method string, url string) (*http.Response, error) {
	var client = &http.Client{}
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
