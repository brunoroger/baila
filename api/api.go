package api

import (
	"baila"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

// temperatureServicer service
type temperatureServicer struct{}

// NewTemperatureServicer return instance temperatureservicer
func NewTemperatureServicer() temperatureServicer {
	return temperatureServicer{}
}

// Temperature type
func (s temperatureServicer) Temperature(city string, unit string) (*baila.Temperature, error) {

	switch unit {
	case "1":
		unit = "metric"
	case "2":
		unit = "imperial"
	default:
		return nil, errors.New("unit not found")
	}

	baseURL, errEncode := url.Parse("https://openweathermap.org/data/2.5/weather")
	if errEncode != nil {
		return nil, errEncode
	}

	params := url.Values{}
	params.Add("q", city)
	params.Add("appid", "b6907d289e10d714a6e88b30761fae22")
	params.Add("units", unit)

	baseURL.RawQuery = params.Encode()

	req, errReq := http.NewRequest("GET", baseURL.String(), nil)
	if errReq != nil {
		return nil, errReq
	}

	res, errDo := http.DefaultClient.Do(req)
	if errDo != nil {
		return nil, errDo
	}

	defer res.Body.Close()
	body, errRead := ioutil.ReadAll(res.Body)
	if errRead != nil {
		return nil, errRead
	}

	var temp baila.Temperature
	errReadyJSON := json.Unmarshal(body, &temp)
	if errReadyJSON != nil {
		return nil, errReadyJSON
	}

	return &temp, nil
}
