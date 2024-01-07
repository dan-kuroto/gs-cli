/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/dan-kuroto/gs-cli/utils"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a gin-stronger application",
	Run: func(cmd *cobra.Command, args []string) {
		projectName := strings.TrimSpace(utils.Input("Project name: "))
		utils.AssertNotEmpty("project name", projectName)
		customConfig := strings.TrimSpace(utils.Input("Use custom configuration? [y/n]: ")) == "y"

		utils.Mkdir(utils.GetPath(projectName))
		utils.Save(utils.GetPath(projectName, "banner.txt"), utils.GetBanner())
		utils.Save(utils.GetPath(projectName, "go.mod"), utils.GetGoMod(projectName))
		utils.Save(utils.GetPath(projectName, "application.yml"), utils.GetApplicationYml(customConfig))
		utils.Save(utils.GetPath(projectName, ".gitignore"), utils.GetGitIgnore())
		utils.Mkdir(utils.GetPath(projectName, "scripts"))
		utils.Save(utils.GetPath(projectName, "scripts", utils.GetScriptName("buildrun")), utils.GetBuildRunScript(projectName))
		utils.Save(utils.GetPath(projectName, "scripts", utils.GetScriptName("build")), utils.GetBuildScript(projectName))
		utils.Save(utils.GetPath(projectName, "scripts", utils.GetScriptName("rundev")), utils.GetRunDevScript(projectName))
		utils.Save(utils.GetPath(projectName, "scripts", utils.GetScriptName("runrelease")), utils.GetRunReleaseScript(projectName))
		utils.Save(utils.GetPath(projectName, projectName+".go"), utils.GetMainGo(projectName, customConfig))
		if customConfig {
			utils.Mkdir(utils.GetPath(projectName, "utils"))
			utils.Save(utils.GetPath(projectName, "utils", "config.go"), utils.GetUtilsConfigGo(projectName))
		}
		utils.Mkdir(utils.GetPath(projectName, "demo"))
		utils.Save(utils.GetPath(projectName, "demo", "demo.go"), utils.GetDemoInitGo(projectName, "demo"))
		utils.Save(utils.GetPath(projectName, "demo", "controller.go"), utils.GetDemoControllerGo(projectName, "demo"))
		utils.Save(utils.GetPath(projectName, "demo", "model.go"), utils.GetDemoModelGo(projectName, "demo"))

		fmt.Print(utils.GetDoneMessage(projectName))
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
