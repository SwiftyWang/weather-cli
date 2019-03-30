package main

import (
	"net/http"
	"io/ioutil"
)

const apiUrl = "http://www.sojson.com/open/api/weather/json.shtml?city="

func Request(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body), nil
}