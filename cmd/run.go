/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/dan-kuroto/gs-cli/utils"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute some shortcut commands",
}

var runBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Compile gin-stronger application",
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.ReadConfig("gs.json")
		config.App.Main = strings.TrimSpace(config.App.Main)
		utils.AssertNotEmpty("app.main in gs.json", config.App.Main)
		config.App.Target = strings.TrimSpace(config.App.Target)
		utils.AssertNotEmpty("app.target in gs.json", config.App.Target)

		fmt.Printf("> go build -o %s %s\n", config.App.Target, config.App.Main)
		command := exec.Command("go", "build", "-o", config.App.Target, config.App.Main)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		if err := command.Run(); err != nil {
			utils.ThrowE(err)
		}
		fmt.Println("Successfully built to", config.App.Target)
	},
}

var runDevCmd = &cobra.Command{
	Use:   "dev",
	Short: "Compile gin-stronger application, and run it in development mode",
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.ReadConfig("gs.json")
		config.App.Main = strings.TrimSpace(config.App.Main)
		utils.AssertNotEmpty("app.main in gs.json", config.App.Main)
		config.App.Target = strings.TrimSpace(config.App.Target)
		utils.AssertNotEmpty("app.target in gs.json", config.App.Target)

		fmt.Printf("> go build -o %s %s\n", config.App.Target, config.App.Main)
		buildCommand := exec.Command("go", "build", "-o", config.App.Target, config.App.Main)
		buildCommand.Stdout = os.Stdout
		buildCommand.Stderr = os.Stderr
		if err := buildCommand.Run(); err != nil {
			utils.ThrowE(err)
		}
		fmt.Println("Successfully built to", config.App.Target)

		fmt.Printf("> %s\n", config.App.Target)
		runCommand := exec.Command(config.App.Target)
		runCommand.Stdout = os.Stdout
		runCommand.Stderr = os.Stderr
		runCommand.Run()
	},
}

func init() {
	runCmd.AddCommand(runBuildCmd)
	runCmd.AddCommand(runDevCmd)

	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
