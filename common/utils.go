package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func InitRequest(URL, token string) (*http.Request, error) {
	uri, err := url.ParseRequestURI(URL)
	if err != nil {
		return nil, err
	}
	h := make(map[string][]string, 1)
	h["authorization"] = []string{fmt.Sprintf("Bearer %s", token)}
	return &http.Request{
		Method: "GET",
		URL:    uri,
		Header: h,
		Close:  true,
	}, nil
}

func FireRequest(request *http.Request) ([]byte, error) {
	do, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer do.Body.Close()
	data, err := ioutil.ReadAll(do.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
