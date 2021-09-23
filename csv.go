package csv

import (
	"encoding/csv"
	"log"
	"os"

	"go.k6.io/k6/js/modules"
)

type CSV struct{}

func init() {
	modules.Register("k6/x/csv", new(CSV))
}

func (c *CSV) Append(path string, data []string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.Write(data)

	if err = w.Error(); err != nil {
		log.Fatal(err)
	}

	w.Flush()
}
