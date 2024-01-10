package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ExecBuild(config Config) {
	config.App.Main = strings.TrimSpace(config.App.Main)
	AssertNotEmpty("app.main in gs.json", config.App.Main)
	config.App.Target = strings.TrimSpace(config.App.Target)
	AssertNotEmpty("app.target in gs.json", config.App.Target)

	fmt.Printf("> go build -o %s %s\n", config.App.Target, config.App.Main)
	command := exec.Command("go", "build", "-o", config.App.Target, config.App.Main)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	if err := command.Run(); err != nil {
		ThrowE(err)
	}
	fmt.Println("Successfully built to", config.App.Target)
}

func ExecRun(config Config) {
	config.App.Main = strings.TrimSpace(config.App.Main)
	AssertNotEmpty("app.main in gs.json", config.App.Main)
	config.App.Target = strings.TrimSpace(config.App.Target)
	AssertNotEmpty("app.target in gs.json", config.App.Target)

	fmt.Printf("> %s\n", config.App.Target)
	runCommand := exec.Command(config.App.Target)
	runCommand.Stdout = os.Stdout
	runCommand.Stderr = os.Stderr
	if err := runCommand.Run(); err != nil {
		ThrowE(err)
	}
}
