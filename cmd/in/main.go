package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/dmechas/concourse-team-resource/concourse"
	"github.com/dmechas/concourse-team-resource/logger"
)

var (
	l logger.Logger
)

func main() {
	logFile, err := ioutil.TempFile("", "concourse-team-resource-in.log")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(os.Stderr, "Logging to %s\n", logFile.Name())

	l = logger.NewLogger(logFile)

	response := concourse.InResponse{}
	err = json.NewEncoder(os.Stdout).Encode(response)
	if err != nil {
		l.Debugf("Exiting with error: %v\n", err)
		log.Fatalln(err)
	}
}
