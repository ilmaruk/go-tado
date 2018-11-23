package tado

import (
	"context"
	"encoding/json"
	"errors"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
)

type TadoConfig struct {
	AuthConfig AuthConfig
	ApiHost    string
}

type AuthConfig struct {
	ClientId         string
	ClientSecret     string
	Scopes           []string
	TokenEndpointURL string
}

var (
	tadoConfig TadoConfig
	oauthConf  *oauth2.Config
	client     *http.Client
	ctx        context.Context
)

func init() {
	tadoConfig = TadoConfig{
		ApiHost: "https://my.tado.com",
		AuthConfig: AuthConfig{
			ClientId:         "tado-web-app",
			ClientSecret:     "wZaRN7rpjn3FoNyF5IFuxg9uMzYJcvOoQ8QWiIqS3hfk6gLhVlG57j5YNoZL2Rtc",
			Scopes:           []string{"home.user"},
			TokenEndpointURL: "https://auth.tado.com/oauth/token",
		},
	}
	ctx = context.Background()
	Config(tadoConfig)
}

func Config(config TadoConfig) {
	tadoConfig = config
	oauthConf = &oauth2.Config{
		ClientID:     tadoConfig.AuthConfig.ClientId,
		ClientSecret: tadoConfig.AuthConfig.ClientSecret,
		Scopes:       tadoConfig.AuthConfig.Scopes,
		Endpoint: oauth2.Endpoint{
			TokenURL: tadoConfig.AuthConfig.TokenEndpointURL,
		},
	}
}

func Authenticate(username, password string) error {
	tok, err := oauthConf.PasswordCredentialsToken(ctx, username, password)
	if err != nil {
		return err
	}

	client = oauthConf.Client(ctx, tok)
	return nil
}

func RunRequest(method, path string, data interface{}) error {
	if client == nil {
		return errors.New("tadogo.Authenticate() not run yet")
	}

	req, err := http.NewRequest(method, tadoConfig.ApiHost+path, nil)
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
