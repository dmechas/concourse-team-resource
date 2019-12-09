package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/dmechas/concourse-team-resource/concourse"
	"github.com/dmechas/concourse-team-resource/fly"
	"github.com/dmechas/concourse-team-resource/logger"
	"github.com/dmechas/concourse-team-resource/out"
)

const (
	flyBinaryName = "fly"
)

var (
	l logger.Logger
)

func main() {
	outDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln(err)
	}
	l = logger.NewLogger(os.Stderr)

	var input concourse.OutRequest

	logFile, err := ioutil.TempFile("", "concourse-team-resource-out.log")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(os.Stderr, "Logging to %s\n", logFile.Name())

	err = json.NewDecoder(os.Stdin).Decode(&input)
	if err != nil {
		fmt.Fprintf(logFile, "Exiting with error: %v\n", err)
		log.Fatalln(err)
	}

	flyBinaryPath := filepath.Join(outDir, flyBinaryName)

	flyCommand := fly.NewCommand(input.Source.Target, l, flyBinaryPath)

	response, err := out.NewCommand(l, flyCommand).Run(input)
	if err != nil {
		l.Debugf("Exiting with error: %v\n", err)
		log.Fatalln(err)
	}

	l.Debugf("Returning output: %+v\n", response)
	err = json.NewEncoder(os.Stdout).Encode(response)
	if err != nil {
		l.Debugf("Exiting with error: %v\n", err)
		log.Fatalln(err)
	}
}
