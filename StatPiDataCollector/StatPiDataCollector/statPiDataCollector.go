package main

import (
	"StatPiDataCollector/onlineReader"
	"fmt"
	//"StatPiDataCollector/localReader"
)

func main() {
	err := onlineReader.GetOnlineWeather("Mosbach")
	if err != nil {
		fmt.Print(err)
	}

	//err = localReader.GetLocalData()
	if err != nil {
		fmt.Print(err)
	}
}
