package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

const Version = "v1.2.2 Victory Knight"
const ShortVersion = "v1.2.2"

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

func GetBanner() string {
	return fmt.Sprintf(`   _____          _____ 
  / ____|        / ____|
 | |  __   ___  | (___  
 | | |_ | |___|  \___ \   こんしや～💫
 | |__| |        ____) |
  \_____|       |_____/   %s
`, Version)
}

func GetGoMod(projectName string) string {
	return fmt.Sprintf(`module %s

go 1.21.0

require (
	github.com/dan-kuroto/gin-stronger %s
	github.com/gin-gonic/gin v1.9.1
)`, projectName, ShortVersion)
}

func GetApplicationYml(customConfig bool) string {
	var builder strings.Builder

	builder.WriteString(`# env:
#   active: dev

gin:
  release: false
  host: 127.0.0.1
  port: 5480
`)
	if customConfig {
		builder.WriteString(`
message: hello world
`)
	}
	builder.WriteString(`
snow-flake:
  start-stmp: 1626779686000
`)

	return builder.String()
}

func GetGitIgnore() string {
	return `# If you prefer the allow list template instead of the deny list, see community template:
# https://github.com/github/gitignore/blob/main/community/Golang/Go.AllowList.gitignore
#
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with ` + "`" + `go test -c` + "`" + `
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/

# Go workspace file
go.work
`
}

func GetDoneMessage(projectName string) string {
	var runCmd string
	if IsWindows() {
		runCmd = "script/buildrun.ps1"
	} else {
		runCmd = "bash script/buildrun.sh"
	}
	return fmt.Sprintf(`
Done. Now run:
  cd %s
  go mod tidy
  %s
`, projectName, runCmd)
}

func GetBuildRunScript(projectName string) string {
	if IsWindows() {
		return fmt.Sprintf(`# build app
go build -o target/%s.exe ./%s.go
if ($LASTEXITCODE -ne 0) {
    exit $LASTEXITCODE
}

# run app
./target/%s.exe
`, projectName, projectName, projectName)
	} else {
		return fmt.Sprintf(`# build app
go build -o target/%s %s.go

code=$?
if [ $code -ne 0 ]; then
    exit $code
fi

# run app
target/%s
`, projectName, projectName, projectName)
	}
}

func GetBuildScript(projectName string) string {
	if IsWindows() {
		return fmt.Sprintf(`# build app
go build -o target/%s.exe ./%s.go
`, projectName, projectName)
	} else {
		return fmt.Sprintf(`# build app
go build -o target/%s %s.go
`, projectName, projectName)
	}
}

func GetRunDevScript(projectName string) string {
	if IsWindows() {
		return fmt.Sprintf(`./target/%s.exe`, projectName)
	} else {
		return fmt.Sprintf(`target/%s`, projectName)
	}
}

func GetRunReleaseScript(projectName string) string {
	if IsWindows() {
		return fmt.Sprintf(`./target/%s.exe --release`, projectName)
	} else {
		return fmt.Sprintf(`target/%s --release`, projectName)
	}
}

func GetMainGo(projectName string, customConfig bool) string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf(`package main

import (
	_ "%s/demo"
`, projectName))
	if customConfig {
		builder.WriteString(fmt.Sprintf(`	"%s/utils"
`, projectName))
	}
	builder.WriteString(`	"net/http"
	"github.com/dan-kuroto/gin-stronger/gs"
	"github.com/gin-gonic/gin"
)

func panicStringHandler(c *gin.Context, err string) {
	c.JSON(http.StatusInternalServerError, gin.H{"err": err})
}

func main() {
	gs.SetGlobalPreffix("/api")
	gs.AddGlobalMiddleware(gs.PackagePanicHandler(panicStringHandler))
`)
	if customConfig {
		builder.WriteString(`	gs.RunApp(&utils.Config)
}
`)
	} else {
		builder.WriteString(`	gs.RunAppDefault()
}
`)
	}

	return builder.String()
}

func GetUtilsConfigGo(projectName string) string {
	return `package utils

import "github.com/dan-kuroto/gin-stronger/gs"

type Configuration struct {
	gs.Configuration ` + "`" + `yaml:",inline"` + "`" + `
	Message          string ` + "`" + `yaml:"message"` + "`" + `
}

var Config Configuration
`
}

func GetDemoInitGo(projectName string, packageName string) string {
	return fmt.Sprintf(`package %s

import "github.com/dan-kuroto/gin-stronger/gs"

func init() {
	gs.UseController(&Controller)
}
`, packageName)
}

func GetDemoControllerGo(projectName string, packageName string) string {
	return fmt.Sprintf(`package %s

import "github.com/dan-kuroto/gin-stronger/gs"

type controller struct{}

var Controller controller

func (*controller) GetRouter() gs.Router {
	return gs.Router{
		Path: "/%s",
		Children: []gs.Router{
			{
				Path:     "/hello",
				Method:   gs.GET | gs.POST,
				Handlers: gs.PackageHandlers(Controller.Hello),
			},
		},
	}
}

func (*controller) Hello(demo *DemoRequst) DemoResponse {
	return DemoResponse{Message: "Hello~" + demo.Name}
}
`, packageName, packageName)
}

func GetDemoModelGo(projectName string, packageName string) string {
	return fmt.Sprintf(`package %s

type DemoRequst struct {
	Name string `+"`"+`form:"name"`+"`"+`
}

type DemoResponse struct {
	Message string `+"`"+`json:"message"`+"`"+`
}
`, packageName)
}
