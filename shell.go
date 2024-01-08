package shell

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"slices"
)

type shellParams string

const (
	NoPanic      shellParams = "noPanic"
	WithoutShell shellParams = "WithoutShell"
	UseStdin     shellParams = "UseStdin"
	UseStdOut    shellParams = "UseStdOut"
)

func execute(command string, input string, params ...shellParams) (string, int, error) {
	logger.Debug("Starting command Execution", "command", command)

	var cmd *exec.Cmd

	if slices.Contains(params, WithoutShell) {
		cmd = exec.Command(command)
	} else {
		cmd = exec.Command("bash", "-c", command)
	}

	// Create a buffer to capture the standard output & error
	var stdoutBuf, stderrBuf bytes.Buffer
	if debug || slices.Contains(params, UseStdOut) {
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
		cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)
	} else {
		cmd.Stdout = &stdoutBuf
		cmd.Stderr = &stderrBuf
	}

	// Handle Input
	if input != "" {
		buffer := bytes.Buffer{}
		buffer.Write([]byte(input))
		cmd.Stdin = &buffer
	} else if slices.Contains(params, UseStdin) {
		cmd.Stdin = os.Stdin
	}

	err := cmd.Run() // Execute the command

	// Handle errors
	if err != nil && slices.Contains(params, NoPanic) {
		if exiterr, ok := err.(*exec.ExitError); ok {
			return "", exiterr.ExitCode(), fmt.Errorf("error executing the command: %s - %v", command, err)
		} else {
			panic(fmt.Errorf("error executing the command: %s - %v", command, err))
		}
	} else if err != nil {
		panic(fmt.Errorf("error executing the command: %s - %v", command, err))
	}

	output := stdoutBuf.String() + stderrBuf.String()
	logger.Debug("Command Completed")

	return output, 0, nil
}

func Sh(command string, params ...shellParams) (string, error) {
	output, _, err := execute(command, "", params...)
	return output, err
}

func ShI(command string, input string, params ...shellParams) (string, error) {
	output, _, err := execute(command, input, params...)
	return output, err
}

func ShA(command string, params ...shellParams) (string, int, error) {
	return execute(command, "", NoPanic)
}
