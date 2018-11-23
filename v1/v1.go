package v1

import (
	"github.com/ilmaruk/tadogo"
	"net/http"
)

const (
	apiPath = "/api/v1"
	mePath = "/me"
)

func Me() (MeStruct, error) {
	var data MeStruct
	err := tado.RunRequest(http.MethodGet, makeUrl(mePath), &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func makeUrl(path string) string {
	return tado.MakeUrl(apiPath + path)
}
