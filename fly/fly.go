package fly

import (
	"bytes"
	"fmt"
	"os/exec"

	"crypto/tls"
	"net/http"

	"github.com/dmechas/concourse-team-resource/logger"
)

//go:generate counterfeiter . Command

type Command interface {
	Login(url string, teamName string, username string, password string, insecure bool) ([]byte, error)
}

type command struct {
	target        string
	logger        logger.Logger
	flyBinaryPath string
}

func NewCommand(target string, logger logger.Logger, flyBinaryPath string) Command {
	return &command{
		target:        target,
		logger:        logger,
		flyBinaryPath: flyBinaryPath,
	}
}

func (f command) Login(
	url string,
	teamName string,
	username string,
	password string,
	insecure bool,
) ([]byte, error) {
	args := []string{
		"login",
		"-c", url,
		"-n", teamName,
	}

	if username != "" && password != "" {
		args = append(args, "-u", username, "-p", password)
	}

	if insecure {
		args = append(args, "-k")
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy:           http.ProxyFromEnvironment,
		}
		http.DefaultClient.Transport = tr
	}

	loginOut, err := f.run(args...)
	if err != nil {
		return nil, err
	}

	syncOut, err := f.run("sync")
	if err != nil {
		return nil, err
	}

	return append(loginOut, syncOut...), nil
}

func (f command) run(args ...string) ([]byte, error) {
	if f.target == "" {
		return nil, fmt.Errorf("target cannot be empty in command.run")
	}

	defaultArgs := []string{
		"-t", f.target,
	}
	allArgs := append(defaultArgs, args...)
	cmd := exec.Command(f.flyBinaryPath, allArgs...)

	outbuf := bytes.NewBuffer(nil)
	errbuf := bytes.NewBuffer(nil)

	cmd.Stdout = outbuf
	cmd.Stderr = errbuf

	f.logger.Debugf("Starting fly command: %v\n", allArgs)
	err := cmd.Start()
	if err != nil {
		// If the command was never started, there will be nothing in the buffers
		return nil, err
	}

	f.logger.Debugf("Waiting for fly command: %v\n", allArgs)
	err = cmd.Wait()
	if err != nil {
		if len(errbuf.Bytes()) > 0 {
			err = fmt.Errorf("%v - %s", err, string(errbuf.Bytes()))
		}
		return outbuf.Bytes(), err
	}

	return outbuf.Bytes(), nil
}
