package tado

import (
	"context"
	"encoding/json"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
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

var authInfo OAuth2Info
var conf = &oauth2.Config{}
var client = &http.Client{}
var ctx = context.Background()

func MakeUrl(path string) string {
	return tadoMyHost + path
}

func Authenticate(userName string, password string) error {
	conf = &oauth2.Config{
		ClientID:     tadoClientId,
		ClientSecret: tadoClientSecret,
		Scopes:       []string{"home.user"},
		Endpoint: oauth2.Endpoint{
			TokenURL: tadoAuthHost + "/oauth/token",
		},
	}

	tok, err := conf.PasswordCredentialsToken(ctx, userName, password)
	if err != nil {
		return err
	}

	client = conf.Client(ctx, tok)
	return nil
}

func NewRequest(method string, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+authInfo.AccessToken)
	return req, nil
}

func RunRequest(method, url string, data interface{}) error {
	req, err := NewRequest(method, url)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, data); err != nil {
		return err
	}

	return nil
}
