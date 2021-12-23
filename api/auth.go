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

const x = `curl 'https://developer.clashroyale.com/api/login' --data-raw '{"email":"savialessiobtc@gmail.com","password":"P#kjN=}H+WFy4B;)"}'`

const y = `curl 'https://developer.clashroyale.com/api/apikey/create' -X POST
-H 'Cookie: session=s%3Aj%3A%7B%22expires%22%3A1640276870087%2C%22uid%22%3A%222bd5a767-03c4-32a3-4f02-839131c77d8d%22%2C%22token%22%3A%22eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJqdGkiOiJlZjY2MDE3MC0zMTdkLThlMDMtYTMwNi0zZTg0MjkyZjJiMzQiLCJzdWIiOiIyYmQ1YTc2Ny0wM2M0LTMyYTMtNGYwMi04MzkxMzFjNzdkOGQiLCJleHAiOjE2NDAyNzY4NzEsImdhbWUiOiJyb3lhbGUiLCJyb2xlIjoiZGV2ZWxvcGVyIn0.R7z6HHP1mImg-yOeJuc5pMQbkDTAkzNoPnAuE-5ITqzxNs4Lq7CIbyXUQ2O39YGBwzMpdkOZIGWBv-5XXY05zA%22%7D.DuY%2FTb6o6jAxVWIskQpIZdwhf8jADXlNoewXegwCQfw; game-api-url=https://api.clashroyale.com/v1/; game-api-token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6ImU4MjUzMTg3LTQ2ODktMThhNC03YWEwLWJlMTc5ZWI2MmMyOSIsImlhdCI6MTY0MDI3MzI3MSwiZXhwIjoxNjQwMjc2ODcxLCJzdWIiOiJkZXZlbG9wZXIvMmJkNWE3NjctMDNjNC0zMmEzLTRmMDItODM5MTMxYzc3ZDhkIiwic2NvcGVzIjpbInJveWFsZSJdLCJsaW1pdHMiOlt7InRpZXIiOiJkZXZlbG9wZXIvYnJvbnplIiwidHlwZSI6InRocm90dGxpbmcifSx7ImNpZHJzIjpbIjk1LjIzOC4xNC41Ni8zMiJdLCJ0eXBlIjoiY2xpZW50In0seyJvcmlnaW5zIjpbImRldmVsb3Blci5jbGFzaHJveWFsZS5jb20iXSwidHlwZSI6ImNvcnMifV19.cD9hKe64v8wMnN7nGAbaw1YjBydb9cf0_kQQaI9g9kEAgFoXZpABX1ZKGBhmlDqM5WMnZBWJ1iY29tVxMQF7EQ' --data-raw '{"name":"key_name","description":"key_description","cidrRanges":["95.238.14.56"],"scopes":null}'`

const z = `curl 'https://developer.clashroyale.com/api/apikey/list' -X POST 
-H 'Cookie: session=s%3Aj%3A%7B%22expires%22%3A1640278599883%2C%22uid%22%3A%222bd5a767-03c4-32a3-4f02-839131c77d8d%22%2C%22token%22%3A%22eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJqdGkiOiIwZjVkYjU2ZS0xYWFlLWEzMTUtZDFlMS00NTI5ZDcwNTdkNTEiLCJzdWIiOiIyYmQ1YTc2Ny0wM2M0LTMyYTMtNGYwMi04MzkxMzFjNzdkOGQiLCJleHAiOjE2NDAyNzg2MDAsImdhbWUiOiJyb3lhbGUiLCJyb2xlIjoiZGV2ZWxvcGVyIn0.TvZqFbUamDBzXoQSf5I8eK5WBx3NIwFMgcQSTG5egWsYt-qKxOo85xzjRLDcLr2FEfHvj5OMr0sUxBirUfJMvw%22%7D.Q6lb8c9BswXwxPqDgCZlwVjDpNurpXWyTs4UjlPtXyI; game-api-url=https://api.clashroyale.com/v1/; game-api-token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6ImE0MDU3MDNhLTRiZDgtNWZlYS0wY2M0LTZmZjZkOWQ4YWUwYSIsImlhdCI6MTY0MDI3NTAwMCwiZXhwIjoxNjQwMjc4NjAwLCJzdWIiOiJkZXZlbG9wZXIvMmJkNWE3NjctMDNjNC0zMmEzLTRmMDItODM5MTMxYzc3ZDhkIiwic2NvcGVzIjpbInJveWFsZSJdLCJsaW1pdHMiOlt7InRpZXIiOiJkZXZlbG9wZXIvYnJvbnplIiwidHlwZSI6InRocm90dGxpbmcifSx7ImNpZHJzIjpbIjk1LjIzOC4xNC41Ni8zMiJdLCJ0eXBlIjoiY2xpZW50In0seyJvcmlnaW5zIjpbImRldmVsb3Blci5jbGFzaHJveWFsZS5jb20iXSwidHlwZSI6ImNvcnMifV19.mPPJTxcxU6zeEXJbXG67x25ZzuGBtO3OaIMsfZEl3m4ean0FXUZAZrcvb_Hxdb_hGbdTGxjjpJ6AkTMbJNzJaw' --data-raw '{}'`

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

const a = `curl 'https://developer.clashroyale.com/api/apikey/revoke' -X POST
-H 'Cookie: session=s%3Aj%3A%7B%22expires%22%3A1640279738613%2C%22uid%22%3A%222bd5a767-03c4-32a3-4f02-839131c77d8d%22%2C%22token%22%3A%22eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJqdGkiOiI0ZGRkZjE1YS01NWMxLTBmOGEtNDFjMy01NzE0ZjUwZTgxZjciLCJzdWIiOiIyYmQ1YTc2Ny0wM2M0LTMyYTMtNGYwMi04MzkxMzFjNzdkOGQiLCJleHAiOjE2NDAyNzk3MzksImdhbWUiOiJyb3lhbGUiLCJyb2xlIjoiZGV2ZWxvcGVyIn0.b3aMPJ0axeskxyZpMvQ5igdcVZhIwqT-pd-CQyk3NiheEZ5CMr9jIOJLfrdlM8LTMiBBLzczmyU3UhThi_1E9Q%22%7D.9HNTsSlhrB0Nuy%2FxPGFgLffCYJz0UTEHyKDZw9DuSu4; game-api-url=https://api.clashroyale.com/v1/; game-api-token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6IjMzYmMxYmY2LTQxOTEtZjE2NC04NTBjLTA2YTliYWE2ODY1MSIsImlhdCI6MTY0MDI3NjEzOSwiZXhwIjoxNjQwMjc5NzM5LCJzdWIiOiJkZXZlbG9wZXIvMmJkNWE3NjctMDNjNC0zMmEzLTRmMDItODM5MTMxYzc3ZDhkIiwic2NvcGVzIjpbInJveWFsZSJdLCJsaW1pdHMiOlt7InRpZXIiOiJkZXZlbG9wZXIvYnJvbnplIiwidHlwZSI6InRocm90dGxpbmcifSx7ImNpZHJzIjpbIjk1LjIzOC4xNC41Ni8zMiJdLCJ0eXBlIjoiY2xpZW50In0seyJvcmlnaW5zIjpbImRldmVsb3Blci5jbGFzaHJveWFsZS5jb20iXSwidHlwZSI6ImNvcnMifV19.jwEhLaHw-j2NF7rwfG_f609yv5Niwl0MymOZ1vZABSaiMhC4j09GZuo36JMHZT-RB6GGyIoutUmPt8d8jo2grA' --data-raw '{"id":"3ceb9bb4-6d0a-4939-b39c-dd632641d983"}'`

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
