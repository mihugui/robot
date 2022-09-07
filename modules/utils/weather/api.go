package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	key = "b1481a75d5624e298eb46bcb57b13d14"
)

type City struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type NowWeather struct {
	Temp      string `json:"temp"`
	Text      string `json:"text"`
	WindDir   string `json:"windDir"`
	FeelsLike string `json:"feelsLike"`
	WindSpeed string `json:"windSpeed"`
}

type Response struct {
	Code  string     `json:"code"`
	Citys []City     `json:"location"`
	Now   NowWeather `json:"now"`
}

func GetWeather(city string) string {

	// 获取城市ID
	cityBean := GetCityID(city)
	res, err := http.Get("https://devapi.qweather.com/v7/weather/now?key=" + key + "&location=" + cityBean.ID)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var response Response

	// string 转json
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	msg := "城市:" + cityBean.Name + "\n" + "温度:" + response.Now.Temp + "℃\n" + "天气:" + response.Now.Text + "\n" + "风向:" + response.Now.WindDir + "\n" + "体感温度:" + response.Now.FeelsLike + "℃\n" + "风速:" + response.Now.WindSpeed + "\n"

	if response.Code == "200" {
		return msg
	} else {
		return ""
	}
}

func GetCityID(city string) City {

	// 获取城市ID
	res, err := http.Get("https://geoapi.qweather.com/v2/city/lookup?key=" + key + "&location=" + url.QueryEscape(city))
	if err != nil {
		fmt.Println(err)
		return City{}
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return City{}
	}
	var response Response

	// string 转json
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return City{}
	}

	if response.Code == "200" {
		return response.Citys[0]
	} else {
		return City{}
	}

}
