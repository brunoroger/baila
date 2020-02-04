package api

import (
	"baila"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	urlPath = "https://openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=%s"
	appid   = "b6907d289e10d714a6e88b30761fae22"
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

	url := fmt.Sprintf(urlPath, city, appid, unit)

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
	errReadyJSON := json.Unmarshal(body, &temp)
	if errReadyJSON != nil {
		return nil, errReadyJSON
	}

	return &temp, nil
}
