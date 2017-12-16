package onlineReader

import (
	"encoding/json"
	"net/http"
	"fmt"
	"time"
	"io/ioutil"
	"StatPiDataCollector/dataWriter"
	"StatPiDataCollector/core"
)


func GetOnlineWeather(city string) (error) {

	apikey := ""
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&units=imperial&APPID=%s", city, apikey)

	httpclient := &http.Client{
		Timeout: time.Second * 2,
	}

	res, err := httpclient.Get(url)
	if err != nil {
		return err
	}

	// defer the closing of the res body
	defer res.Body.Close()

	// read the http response body into a byte stream
	bodybyte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	data := core.CurrentWeatherResponse{}

	// unmarshal the byte stream into a Go data type
	err = json.Unmarshal(bodybyte, &data)
	if err != nil {
		return  err
	}

    	temp :=(data.Temp -32) / 1.8
    	fmt.Print(temp)

	writedata := core.New(temp, data.Humidity, int64(data.Dt), "online")
    	err = dataWriter.WriteData(writedata, "data.csv")
	if err != nil {
		return err
	}
	return nil

}
