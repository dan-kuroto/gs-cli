package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ExecBuild(config *Config) *exec.Cmd {
	config.App.Main = strings.TrimSpace(config.App.Main)
	AssertNotEmpty("app.main in gs.json", config.App.Main)
	config.App.Target = strings.TrimSpace(config.App.Target)
	AssertNotEmpty("app.target in gs.json", config.App.Target)

	fmt.Printf("> go build -o %s %s\n", config.App.Target, config.App.Main)
	command := exec.Command("go", "build", "-o", config.App.Target, config.App.Main)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Start(); err != nil {
		ThrowE(err)
	}

	return command
}

func WaitExecBuild(config *Config) {
	if err := ExecBuild(config).Wait(); err != nil {
		ThrowE(err)
	}
	fmt.Println("Successfully built to", config.App.Target)
}

func ExecRun(config *Config) *exec.Cmd {
	config.App.Main = strings.TrimSpace(config.App.Main)
	AssertNotEmpty("app.main in gs.json", config.App.Main)
	config.App.Target = strings.TrimSpace(config.App.Target)
	AssertNotEmpty("app.target in gs.json", config.App.Target)

	fmt.Printf("> %s\n", config.App.Target)
	command := exec.Command(config.App.Target)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Start(); err != nil {
		ThrowE(err)
	}

	return command
}

func WaitExecRun(config *Config) {
	if err := ExecRun(config).Wait(); err != nil {
		ThrowE(err)
	}
}

func KillProcess(command *exec.Cmd) {
	if err := command.Process.Kill(); err != nil {
		ThrowE(err)
	}
}
