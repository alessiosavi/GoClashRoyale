package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	stringutils "github.com/alessiosavi/GoGPUtils/string"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Keys struct {
	Status struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Detail  interface{} `json:"detail"`
	} `json:"status"`
	SessionExpiresInSeconds int `json:"sessionExpiresInSeconds"`
	Keys                    []struct {
		ID          string      `json:"id"`
		DeveloperID string      `json:"developerId"`
		Tier        string      `json:"tier"`
		Name        string      `json:"name"`
		Description string      `json:"description"`
		Origins     interface{} `json:"origins"`
		Scopes      []string    `json:"scopes"`
		CidrRanges  []string    `json:"cidrRanges"`
		ValidUntil  interface{} `json:"validUntil"`
		Key         string      `json:"key"`
	} `json:"keys"`
}

func getIp() (string, error) {
	get, err := http.Get("https://ifconfig.co")
	if err != nil {
		return "", err
	}
	defer get.Body.Close()
	all, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return "", err
	}
	return string(all), nil
}

func NewKey(email, password string) (string, error) {
	cookies, err := Login(email, password)
	if err != nil {
		return "", err
	}

	keys, cookies, err := ListKeys(cookies)
	if err != nil {
		return "", err
	}
	if len(keys.Keys) == 0 {
		cookies, err = GenerateApiKey(cookies)
		if err != nil {
			return "", err
		}
		keys, cookies, err = ListKeys(cookies)
		if err != nil {
			return "", err
		}
	}
	return keys.Keys[len(keys.Keys)-1].Key, nil

}

func ListKeys(cookies []*http.Cookie) (*Keys, []*http.Cookie, error) {
	// List the created API KEY
	request, err := http.NewRequest("POST", "https://developer.clashroyale.com/api/apikey/list", bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		return nil, nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	for _, cookie := range cookies {
		request.AddCookie(cookie)
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	var keys Keys
	if err = json.Unmarshal(all, &keys); err != nil {
		return nil, nil, err
	}
	for _, cookie := range cookies {
		cookies = append(cookies, cookie)
	}
	return &keys, cookies, nil
}

func GenerateApiKey(cookies []*http.Cookie) ([]*http.Cookie, error) {
	// Generate a new API KEY
	ip, err := getIp()
	if err != nil {
		return nil, err
	}
	ip = stringutils.Trim(ip)

	body := []byte(`{"name":"GoClashRoyaleKey","description":"GoClashRoyaleKey","cidrRanges":["` + ip + `"],"scopes":null}`)
	request, err := http.NewRequest("POST", "https://developer.clashroyale.com/api/apikey/create", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	for _, cookie := range cookies {
		request.AddCookie(cookie)
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	_, _ = io.Copy(ioutil.Discard, resp.Body)
	for _, cookie := range resp.Cookies() {
		cookies = append(cookies, cookie)
	}
	return cookies, nil
}

func Login(email string, password string) ([]*http.Cookie, error) {
	// Login and get the cookie
	data, err := http.Post("https://developer.clashroyale.com/api/login", "application/json; charset=utf-8", strings.NewReader(fmt.Sprintf(`{"email":"%s","password":"%s"}`, email, password)))
	if err != nil {
		return nil, err
	}
	return data.Cookies(), nil
}

func DeleteKey(cookies []*http.Cookie, keyName string) ([]*http.Cookie, error) {
	// Delete an API KEY by the given keyName
	keys, cookies, err := ListKeys(cookies)
	if err != nil {
		return nil, err
	}
	for _, key := range keys.Keys {
		if key.Name == keyName {
			request, err := http.NewRequest("POST", "https://developer.clashroyale.com/api/apikey/revoke", bytes.NewBuffer([]byte(`{"id":"`+key.ID+`"}`)))
			request.Header.Set("Content-Type", "application/json")
			for _, cookie := range cookies {
				request.AddCookie(cookie)
			}

			resp, err := http.DefaultClient.Do(request)
			if err != nil {
				return nil, err
			}

			_, _ = io.Copy(ioutil.Discard, resp.Body)
			_ = resp.Body.Close()
			for _, cookie := range resp.Cookies() {
				cookies = append(cookies, cookie)
			}
		}

	}
	return cookies, nil
}
