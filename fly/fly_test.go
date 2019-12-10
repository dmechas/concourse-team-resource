package fly_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/dmechas/concourse-team-resource/fly"
	"github.com/dmechas/concourse-team-resource/logger"
	"github.com/stretchr/testify/assert"
)

var (
	target          string
	flyBinaryPath   string
	teamName        string
	localUser       string
	githubTeam      string
	tempDir         string
	fakeFlyContents string
)

func TestGetTeam(t *testing.T) {
	target = "some-target"
	fakeLogger := logger.NewLogger(os.Stderr)
	teamName = "some-team"
	var err error
	tempDir, err = ioutil.TempDir("", "")
	assert.NoError(t, err)

	flyBinaryPath = filepath.Join(tempDir, "fake_fly")
	fakeFlyContents = `#!/bin/bash
	echo -n "$@"`

	err = ioutil.WriteFile(flyBinaryPath, []byte(fakeFlyContents), os.ModePerm)
	assert.NoError(t, err)

	flyCommand := fly.NewCommand(target, fakeLogger, flyBinaryPath)

	output, err := flyCommand.GetTeam(teamName)
	assert.NoError(t, err)

	expected := fmt.Sprintf("%s %s %s %s %s",
		"-t", target, "get-team", "-n", teamName)
	assert.Equal(t, expected, string(output))
}

func TestSetTeam(t *testing.T) {
	target = "some-target"
	fakeLogger := logger.NewLogger(os.Stderr)
	teamName = "some-team"
	localUser = "some-user"
	githubTeam = "some-team"
	var err error
	tempDir, err = ioutil.TempDir("", "")
	assert.NoError(t, err)

	flyBinaryPath = filepath.Join(tempDir, "fake_fly")
	fakeFlyContents = `#!/bin/bash
	echo -n "$@"`

	err = ioutil.WriteFile(flyBinaryPath, []byte(fakeFlyContents), os.ModePerm)
	assert.NoError(t, err)

	flyCommand := fly.NewCommand(target, fakeLogger, flyBinaryPath)

	output, err := flyCommand.SetTeam(teamName, localUser, githubTeam)
	assert.NoError(t, err)

	expected := fmt.Sprintf("%s %s %s %s %s %s %s=%s %s=%s",
		"-t", target,
		"set-team",
		"--non-interactive",
		"-n", teamName,
		"--local-user", localUser,
		"--github-team", githubTeam,
	)
	assert.Equal(t, expected, string(output))
}
