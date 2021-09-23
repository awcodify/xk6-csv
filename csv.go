package csv

import (
	"encoding/csv"
	"os"

	"go.k6.io/k6/js/modules"
)

type CSV struct{}

func init() {
	modules.Register("k6/x/csv", new(CSV))
}

func (c *CSV) Append(path string, data []string) (err error) {
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
