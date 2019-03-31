package main

type Response struct {
	Time     string `json:"time"`
	CityInfo struct {
		City       string `json:"city"`
		CityID     string `json:"cityId"`
		Parent     string `json:"parent"`
		UpdateTime string `json:"updateTime"`
	} `json:"cityInfo"`
	Date    string `json:"date"`
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    struct {
		Shidu     string  `json:"shidu"`
		Pm25      float64 `json:"pm25"`
		Pm10      float64 `json:"pm10"`
		Quality   string  `json:"quality"`
		Wendu     string  `json:"wendu"`
		Ganmao    string  `json:"ganmao"`
		Yesterday struct {
			Date    string  `json:"date"`
			Sunrise string  `json:"sunrise"`
			High    string  `json:"high"`
			Low     string  `json:"low"`
			Sunset  string  `json:"sunset"`
			Aqi     float64 `json:"aqi"`
			Ymd     string  `json:"ymd"`
			Week    string  `json:"week"`
			Fx      string  `json:"fx"`
			Fl      string  `json:"fl"`
			Type    string  `json:"type"`
			Notice  string  `json:"notice"`
		} `json:"yesterday"`
		Forecast []struct {
			Date    string  `json:"date"`
			Sunrise string  `json:"sunrise"`
			High    string  `json:"high"`
			Low     string  `json:"low"`
			Sunset  string  `json:"sunset"`
			Aqi     float64 `json:"aqi,omitempty"`
			Ymd     string  `json:"ymd"`
			Week    string  `json:"week"`
			Fx      string  `json:"fx"`
			Fl      string  `json:"fl"`
			Type    string  `json:"type"`
			Notice  string  `json:"notice"`
		} `json:"forecast"`
	} `json:"data"`
}

type Cities []struct {
	ID       int    `json:"_id"`
	ID2      int    `json:"id"`
	Pid      int    `json:"pid"`
	CityCode string `json:"city_code"`
	CityName string `json:"city_name"`
}
