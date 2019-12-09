package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/dmechas/concourse-team-resource/concourse"
	"github.com/dmechas/concourse-team-resource/logger"
)

var (
	l logger.Logger
)

func main() {
	l = logger.NewLogger()

	response := concourse.InResponse{}
	err := json.NewEncoder(os.Stdout).Encode(response)
	if err != nil {
		l.Debugf("Exiting with error: %v\n", err)
		log.Fatalln(err)
	}
}
