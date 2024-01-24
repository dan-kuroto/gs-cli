package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"text/template"

	"github.com/dan-kuroto/gs-cli/tpl"
)

const Version = "v1.2.3 Victory Knight"
const ShortVersion = "v1.2.3"

func execTemplate(tplStr string, data any) string {
	tmpl, err := template.New("").Parse(tplStr)
	if err != nil {
		ThrowE(err)
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, data)
	if err != nil {
		ThrowE(err)
	}
	return buffer.String()
}

func GetGSJson(projectName string, customConfig bool) string {
	var target string
	if IsWindows() {
		target = fmt.Sprintf("./target/%s.exe", projectName)
	} else {
		target = fmt.Sprintf("target/%s", projectName)
	}
	jsonData, err := json.MarshalIndent(Config{
		GsVersion: ShortVersion,
		App: AppConfig{
			Name:         projectName,
			Version:      "0.0.0",
			CustomConfig: customConfig,
			Main:         fmt.Sprintf(`%s.go`, projectName),
			Target:       target,
		},
	}, "", "  ")
	if err != nil {
		ThrowE(err)
	}
	return string(jsonData)
}

func GetBannerTxt() string {
	return execTemplate(tpl.BannerTxt, map[string]any{
		"version": Version,
	})
}

func GetGoMod(projectName string) string {
	return execTemplate(tpl.GoMod, map[string]any{
		"projectName":  projectName,
		"shortVersion": ShortVersion,
	})
}

func GetApplicationYml(customConfig bool) string {
	return execTemplate(tpl.ApplicationYml, map[string]any{
		"customConfig": customConfig,
	})
}

func GetGitIgnore() string {
	return execTemplate(tpl.GitIgnore, map[string]any{})
}

func GetCreateDoneMessage(projectName string) string {
	return fmt.Sprintf(`
Done. Now run:
  cd %s
  go mod tidy
  gs-cli run dev
`, projectName)
}

func GetMainGo(projectName string, customConfig bool) string {
	return execTemplate(tpl.MainGo, map[string]any{
		"projectName":  projectName,
		"customConfig": customConfig,
	})
}

func GetUtilsConfigGo(projectName string) string {
	return execTemplate(tpl.UtilsConfigGo, map[string]any{})
}

func GetDemoDemoGo(projectName string, packageName string) string {
	return execTemplate(tpl.DemoDemoGo, map[string]any{
		"packageName": packageName,
	})
}

func GetDemoControllerGo(projectName string, packageName string) string {
	return execTemplate(tpl.DemoControllerGo, map[string]any{
		"packageName": packageName,
	})
}

func GetDemoModelGo(projectName string, packageName string) string {
	return execTemplate(tpl.DemoModelGo, map[string]any{
		"packageName": packageName,
	})
}
