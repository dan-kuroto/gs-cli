/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/dan-kuroto/gs-cli/utils"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a gin-stronger application",
	Run: func(cmd *cobra.Command, args []string) {
		projectName := utils.Input("Project name: ")
		customConfig := utils.Input("Use custom configuration? [y/n]: ") == "y"

		if err := os.Mkdir(utils.GetPath(projectName), fs.ModeDir); err != nil {
			panic(err)
		}
		utils.Save(utils.GetPath(projectName, "bannner.txt"), utils.GetBanner())
		utils.Save(utils.GetPath(projectName, "go.mod"), utils.GetGoMod(projectName))
		utils.Save(utils.GetPath(projectName, "application.yml"), utils.GetApplicationYml(customConfig))
		utils.Save(utils.GetPath(projectName, ".gitignore"), utils.GetGitIgnore())

		// TODO: 获取系统类型并生成scripts文件路径
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
