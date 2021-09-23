package redis

import (
	"encoding/csv"
	"log"
	"os"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/csv", new(Redis))
}

func append(path string, data string) [][]string {
	f, err := os.Open(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var data [][]string

	w := csv.NewWriter(f)
	w.Write((data))

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
