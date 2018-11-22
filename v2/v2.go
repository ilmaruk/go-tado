package v2

import (
	"encoding/json"
	"fmt"
	"github.com/ilmaruk/tadogo"
	"io/ioutil"
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

func GetZones(houseId int) ([]Zone, error) {
	var data []Zone

	req, err := tado.NewRequest(http.MethodGet, makeUrl(fmt.Sprintf(zonesPath, houseId)))
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

func GetZoneState(houseId int, zoneId int) (ZoneState, error) {
	var data ZoneState

	req, err := tado.NewRequest(http.MethodGet, makeUrl(fmt.Sprintf(zoneStatePath, houseId, zoneId)))
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
