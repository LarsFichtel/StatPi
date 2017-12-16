package dataWriter

import (
	"StatPiDataCollector/core"
	"os"
	"encoding/csv"
)
func WriteData (data core.Entry, file string) error{
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	dataslice := core.EntryToString(data)
	w := csv.NewWriter(f)
	w.UseCRLF = true //add newline
	w.Write(dataslice )
	w.Flush()
	return nil
}