package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiUrl = "http://t.weather.sojson.com/api/weather/city/"

func readJsonFile() (Cities, error) {
	bytes, _ := ioutil.ReadFile("_city.json")
	var cities Cities
	error := json.Unmarshal(bytes, &cities)
	return cities, error
}

func Request(city string) (string, error) {
	cities, e := readJsonFile()
	if e != nil {
		return "", e
	}
	var code string
	for _, c := range cities {
		if c.CityName == city {
			code = c.CityCode
			break
		}
	}
	if len(code) == 0 {
		return "", fmt.Errorf("cannot find city code by given city:%s", city)
	}
	url := apiUrl + code
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return string(body), nil
}
