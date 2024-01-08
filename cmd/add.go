/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"

	"github.com/dan-kuroto/gs-cli/utils"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new package in gin-stronger application",
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.ReadConfig("gs.json")
		config.App.Name = strings.TrimSpace(config.App.Name)
		utils.AssertNotEmpty("app.name in gs.json", config.App.Name)
		config.App.Main = strings.TrimSpace(config.App.Main)
		utils.AssertNotEmpty("app.main in gs.json", config.App.Main)

		projectName := config.App.Name
		mainPath := utils.GetPath(config.App.Main)
		packageName := strings.TrimSpace(utils.Input("Package name: "))
		utils.AssertNotEmpty("package name", packageName)

		utils.Mkdir(utils.GetPath(packageName))
		utils.Save(utils.GetPath(packageName, packageName+".go"), utils.GetDemoInitGo(projectName, packageName))
		utils.Save(utils.GetPath(packageName, "controller.go"), utils.GetDemoControllerGo(projectName, packageName))
		utils.Save(utils.GetPath(packageName, "model.go"), utils.GetDemoModelGo(projectName, packageName))

		utils.AddPackageToMainGo(projectName, mainPath, packageName)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
