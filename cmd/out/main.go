package main

import (
	"fmt"
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
	if len(os.Args) < 2 {
		log.Fatalln(fmt.Sprintf(
			"not enough args - usage: %s <sources directory>", os.Args[0]))
	}

	sourcesDir := os.Args[1]

	outDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln(err)
	}

	var input concourse.OutRequest

	flyBinaryPath := filepath.Join(outDir, flyBinaryName)

	l = logger.NewLogger()

	input.Source.Target = "AHAHAH"

	flyCommand := fly.NewCommand(input.Source.Target, l, flyBinaryPath)

	response, err := out.NewCommand(l, flyCommand, sourcesDir).Run(input)
	if err != nil {
		l.Debugf("Exiting with error: %v\n", err)
		log.Fatalln(err)
	}

	l.Debugf("Returning output: %+v\n", response)

}
