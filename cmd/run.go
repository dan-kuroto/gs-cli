/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

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
	Short: "Compile application",
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.ReadConfig("gs.json")

		utils.WaitExecBuild(&config)
	},
}

var runDevCmd = &cobra.Command{
	Use:   "dev",
	Short: "Compile application and run in dev mode",
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.ReadConfig("gs.json")

		utils.WaitExecBuild(&config)
		utils.WaitExecRun(&config)
	},
}

var runWatchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Compile application and run in dev mode, then automatically redo when code is modified",
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.ReadConfig("gs.json")

		modified := make(chan bool, 1)
		utils.WaitExecBuild(&config)
		command := utils.ExecRun(&config)
		for {
			select {
			case <-modified:
				fmt.Println()
				fmt.Println("File modification detected. Kill current process ...")

				utils.KillProcess(command)
				command.Wait()

				fmt.Println("Do recompile & restart ...")
				fmt.Println()

				utils.WaitExecBuild(&config)
				command = utils.ExecRun(&config)
			case <-time.After(500 * time.Millisecond):
				if utils.CheckModified() {
					modified <- true
				}
			}
		}
	},
}

func init() {
	runCmd.AddCommand(runBuildCmd)
	runCmd.AddCommand(runDevCmd)
	runCmd.AddCommand(runWatchCmd)

	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
