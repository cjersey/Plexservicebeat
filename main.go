package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/cjersey/plexservicebeat/beater"
)

func main() {
	err := beat.Run("plexservicebeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
