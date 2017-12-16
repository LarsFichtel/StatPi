package localReader

import (
	"io/ioutil"
	"StatPiDataCollector/dataWriter"
	"StatPiDataCollector/core"
	"strconv"
	"time"
)

func GetLocalData()error{
	tempfile := "data.csv" //filetobeset
	temp, err := ioutil.ReadFile(tempfile)
	if err != nil {
		return err
	}
	tempstring  := string(temp)

	humfile := "data.csv" //filetobeset
	hum, err := ioutil.ReadFile(humfile)
	if err != nil {
		return err
	}
	humstring  := string(hum)

	tempfloat, err := strconv.ParseFloat(tempstring, 64)
	if err != nil {
		return err
	}

	humfloat, err := strconv.Atoi(humstring)
	if err != nil {
		return err
	}

	writedata := core.New(tempfloat, humfloat, time.Now().Unix(), "online")
	err = dataWriter.WriteData(writedata, "data.csv")
	if err != nil {
		return err
	}
	return nil
}