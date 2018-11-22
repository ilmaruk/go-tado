package v2

import "time"

type Home struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DeviceMetadata struct {
	Locale    string `json:"locale"`
	Model     string `json:"model"`
	OsVersion string `json:"osVersion"`
	Platform  string `json:"Platform"`
}

type DeviceBearingFromHome struct {
	Degrees float64 `json:"degrees"`
	Radians float64 `json:"radians"`
}

type DeviceLocation struct {
	AtHome                        bool                  `json:"atHome"`
	BearingFromHome               DeviceBearingFromHome `json:"bearingFromHome"`
	RelativeDistanceFromHomeFence float64               `json:"relativeDistanceFromHomeFence"`
	Stale                         bool                  `json:"stale"`
}

type DeviceSettings struct {
	GeoTrackingEnabled bool `json:"geoTrackingEnabled"`
}

type MobileDevice struct {
	Metadata DeviceMetadata `json:"deviceMetadata"`
	Id       int            `json:"id"`
	Location DeviceLocation `json:"location"`
	Name     string         `json:"name"`
	Settings DeviceSettings `json:"settings"`
}

type Me struct {
	Email         string         `json:"email"`
	Homes         []Home         `json:"homes"`
	Id            string         `json:"id"`
	Locale        string         `json:"locale"`
	MobileDevices []MobileDevice `json:"mobileDevices"`
	Name          string         `json:"name"`
	Username      string         `json:"username"`
}

type Zone struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type HumidityDataPoint struct {
	Percentage float32   `json:"percentage"`
	Timestamp  time.Time `json:"timestamp"`
	Type       string    `json:"type"`
}

type temperaturePrecision struct {
	Celsius    float32 `json:"celsius"`
	Fahrenheit float32 `json:"fahrenheit"`
}

type InsideTemperatureDataPoint struct {
	Celsius    float32              `json:"celsius"`
	Fahrenheit float32              `json:"fahrenheit"`
	Precision  temperaturePrecision `json:"precision"`
	Timestamp  time.Time            `json:"timestamp"`
	Type       string               `json:"type"`
}

type ZoneSensorDataPoints struct {
	Humidity          HumidityDataPoint          `json:"humidity"`
	InsideTemperature InsideTemperatureDataPoint `json:"insideTemperature"`
}

type ZoneState struct {
	SensorDataPoints ZoneSensorDataPoints `json:"sensorDataPoints"`
	TadoMode         string               `json:"tadoMode"`
}
