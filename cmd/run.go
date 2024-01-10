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
		utils.ExecBuild(config)
	},
}

var runDevCmd = &cobra.Command{
	Use:   "dev",
	Short: "Compile application and run in dev mode",
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.ReadConfig("gs.json")
		utils.ExecBuild(config)
		utils.ExecRun(config)
	},
}

var runWatchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Compile application and run in dev mode, then automatically redo when code is modified",
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.ReadConfig("gs.json")
		go func() {
			<-time.After(time.Second * 10)
			fmt.Println("超时 ~")
			// NOTE: 不过用os.Exit退出的话子进程关不掉
			utils.ThrowS("主动结束")
		}()
		utils.ExecBuild(config)
		utils.ExecRun(config)
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
