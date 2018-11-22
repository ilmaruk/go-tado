package tado

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	tadoClientId     = "tado-web-app"
	tadoClientSecret = "wZaRN7rpjn3FoNyF5IFuxg9uMzYJcvOoQ8QWiIqS3hfk6gLhVlG57j5YNoZL2Rtc"
	tadoAuthHost     = "https://auth.tado.com"
	tadoMyHost       = "https://my.tado.com"
)

type OAuth2Info struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	Jti          string `json:"jti"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

var (
	HttpClient = &http.Client{}
	authInfo   OAuth2Info
)

func MakeUrl(path string) string {
	return tadoMyHost + path
}

func Authenticate(userName string, password string) error {
	data := url.Values{
		"client_id":     {tadoClientId},
		"client_secret": {tadoClientSecret},
		"grant_type":    {"password"},
		"scope":         {"home.user"},
		"password":      {password},
		"username":      {userName},
	}
	req, err := http.NewRequest(http.MethodPost, tadoAuthHost+ "/oauth/token", strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := HttpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &authInfo); err != nil {
		return err
	}

	return nil
}

func NewRequest(method string, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer " + authInfo.AccessToken)
	return req, nil
}