package csv

import (
	"encoding/csv"
	"os"
)

func Append(path string, data [][]string) error {
	f, err := os.Open(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.Write(data)

	if err := w.Error(); err != nil {
		return
	}

	return
}
