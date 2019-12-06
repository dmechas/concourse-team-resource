package out

import (
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
	sourcesDir string,
) *Command {
	return &Command{
		logger:     logger,
		flyCommand: flyCommand,
		sourcesDir: sourcesDir,
	}
}

func (c *Command) Run(input concourse.OutRequest) (concourse.OutResponse, error) {
	c.logger.Debugf("Received input: %+v\n", input)
	response := concourse.OutResponse{}

	return response, nil
}
