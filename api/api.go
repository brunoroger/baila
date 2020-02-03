package api

import (
	"baila"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// TemperatureServicer service
type TemperatureServicer struct{}

// Temperature type
func (s TemperatureServicer) Temperature(city string, unit string) (*baila.Temperature, error) {

	switch unit {
	case "1":
		unit = "metric"
	case "2":
		unit = "imperial"
	default:
		return nil, errors.New("unit not found")
	}

	url := "https://openweathermap.org/data/2.5/weather?q=" + city + "&appid=b6907d289e10d714a6e88b30761fae22&units=" + unit

	req, errReq := http.NewRequest("GET", url, nil)
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
	json.Unmarshal(body, &temp)

	return &temp, nil
}
