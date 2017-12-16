package core

import (
	"strconv"
	"fmt"
)

type Entry struct {
	temperature       float64
	humidity          int
	CreationTime      int64 //Unix Time Stamp
	source            string
}

//Example from https://openweathermap.org/current
/*
{"coord":{"lon":139,"lat":35},
"sys":{"country":"JP","sunrise":1369769524,"sunset":1369821049},
"weather":[{"id":804,"main":"clouds","description":"overcast clouds","icon":"04n"}],
"main":{"temp":289.5,"humidity":89,"pressure":1013,"temp_min":287.04,"temp_max":292.04},
"wind":{"speed":7.31,"deg":187.002},
"rain":{"3h":0},
"clouds":{"all":92},
"dt":1369824698,
"id":1851632,
"name":"Shuzenji",
"cod":200}
 */

//struct to help get the CurrentWeatherResponse
type City struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Coordinates struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type CurrentWeather struct {
	Id int `json:"id"`
	Main string `json:"main"`
	Description string `json:"description"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg float64 `json:"deg"`
}

type Clouds struct {
	All int `json:"all"`
}

type Rain struct {
	Threehr int `json:"3h"`
}

type Main struct {
	Temp float64 `json:"temp"`
	Pressure int `json:"pressure"`
	Humidity int `json:"humidity"`
	Temp_min float64 `json:"temp_min"`
	Temp_max float64 `json:"temp_max"`
}
// type of openweathermap response
type CurrentWeatherResponse struct {
	Coordinates `json:"coord"`
	CurrentWeather []CurrentWeather `json:"weather"`
	Main `json:"main"`
	Wind `json:"wind"`
	Rain `json:"rain"`
	Clouds `json:"clouds"`
	Dt int `json:"dt"`
	Id int `json:"id"`
	Name string `json:"name"`
}

func New(temp float64, hum int, time int64, source string) Entry {
	entry := Entry{temp, hum, time, source}
	return entry
}
func EntryToString(entry Entry) []string{
	time64 :=  fmt.Sprintf("%v", entry.CreationTime)
	return []string{strconv.FormatFloat(entry.temperature, 'f', 6, 64), strconv.Itoa(entry.humidity), time64, entry.source}
	//return strconv.FormatFloat(entry.temperature, 'f', 6, 64) + comma + strconv.Itoa(entry.humidity) + comma + strconv.Itoa(entry.CreationTime) + comma + entry.source
}