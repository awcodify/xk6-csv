package csv

import (
	"encoding/csv"
	"os"
)

type CSV struct{}

func Append(path string, data []string) (err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.Write(data)

	if err = w.Error(); err != nil {
		return
	}

	return
}
