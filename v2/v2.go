package v2

import (
	"fmt"
	"github.com/ilmaruk/tadogo"
	"net/http"
)

const (
	apiPath       = "/api/v2"
	mePath        = "/me"
	zonesPath     = "/homes/%d/zones"
	zoneStatePath = "/homes/%d/zones/%d/state"
)

func GetMe() (Me, error) {
	var data Me
	err := tado.RunRequest(http.MethodGet, apiPath+mePath, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func GetZones(houseId int) ([]Zone, error) {
	var data []Zone
	err := tado.RunRequest(http.MethodGet, apiPath+fmt.Sprintf(zonesPath, houseId), &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func GetZoneState(houseId int, zoneId int) (ZoneState, error) {
	var data ZoneState
	err := tado.RunRequest(http.MethodGet, apiPath+fmt.Sprintf(zoneStatePath, houseId, zoneId), &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
