package v1

import (
	"github.com/ilmaruk/tadogo"
	"net/http"
)

const (
	apiPath = "/api/v1"
	mePath  = "/me"
)

var requestRunner = tado.RunRequest

func GetMe() (Me, error) {
	var data Me
	err := requestRunner(http.MethodGet, apiPath+mePath, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
