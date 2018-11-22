package v1

import (
	"encoding/json"
	"github.com/ilmaruk/tadogo"
	"io/ioutil"
	"net/http"
)

const (
	apiPath = "/api/v1"
	mePath = "/me"
)

func Me() (MeStruct, error) {
	var data MeStruct

	req, err := tado.NewRequest(http.MethodGet, makeUrl(mePath))
	if err != nil {
		return data, err
	}

	resp, err := tado.HttpClient.Do(req)
	if err != nil {
		return data, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &data); err != nil {
		return data, err
	}

	return data, nil
}

func makeUrl(path string) string {
	return tado.MakeUrl(apiPath + path)
}
