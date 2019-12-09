package out

import (
	"crypto/md5"
	"fmt"
	"os"
	"strconv"

	"github.com/dmechas/concourse-team-resource/concourse"
	"github.com/dmechas/concourse-team-resource/fly"
	"github.com/dmechas/concourse-team-resource/logger"
)

const (
	apiPrefix = "/api/v1"
)

type Command struct {
	logger     logger.Logger
	flyCommand fly.Command
	sourcesDir string
}

func NewCommand(
	logger logger.Logger,
	flyCommand fly.Command,
) *Command {
	return &Command{
		logger:     logger,
		flyCommand: flyCommand,
	}
}

func (c *Command) Run(input concourse.OutRequest) (concourse.OutResponse, error) {
	c.logger.Debugf("Received input: %+v\n", input)
	insecure := false
	if input.Source.Insecure != "" {
		var err error
		insecure, err = strconv.ParseBool(input.Source.Insecure)
		if err != nil {
			return concourse.OutResponse{}, err
		}
	}

	c.logger.Debugf("Performing login\n")
	_, err := c.flyCommand.Login(
		input.Source.Target,
		input.Source.MainTeam,
		input.Source.Username,
		input.Source.Password,
		insecure,
	)
	if err != nil {
		return concourse.OutResponse{}, err
	}

	c.logger.Debugf("Login successful\n")

	setOutput, err := c.flyCommand.SetTeam(input.Params.TeamName, input.Params.LocalUser, input.Params.GitHubTeam)
	c.logger.Debugf("pipeline '%s' set; output:\n\n%s\n", input.Params.TeamName, string(setOutput))
	fmt.Fprintf(os.Stderr, "pipeline '%s' set; output:\n\n%s\n", input.Params.TeamName, string(setOutput))
	if err != nil {
		return concourse.OutResponse{}, err
	}

	outBytes, err := c.flyCommand.GetTeam(
		input.Params.TeamName,
	)
	if err != nil {
		return concourse.OutResponse{}, err
	}

	version := fmt.Sprintf(
		"%x",
		md5.Sum(outBytes),
	)

	response := concourse.OutResponse{
		Version: concourse.Version{
			input.Params.TeamName: version,
		},
		Metadata: []concourse.Metadata{},
	}
	return response, nil
}
