package utils

import (
	"fmt"
	"strings"
)

const Version = "v1.1.0 Ginga Strium"

func GetBanner() string {
	return fmt.Sprintf(`   _____          _____ 
  / ____|        / ____|
 | |  __   ___  | (___  
 | | |_ | |___|  \___ \   こんしや～💫
 | |__| |        ____) |
  \_____|       |_____/   %s
`, Version)
}

func GetGoMod(packageName string) string {
	return fmt.Sprintf(`module %s

go 1.21.0

require (
	github.com/dan-kuroto/gin-stronger v1.1.0
	github.com/gin-gonic/gin v1.9.1
)`, packageName)
}

func GetApplicationYml(customConfig bool) string {
	var builder strings.Builder

	builder.WriteString(`# env:
#   active: dev

gin:
  release: true
  host: 127.0.0.1
  port: 5480
`)
	if customConfig {
		builder.WriteString(`
message: hello world
`)
	}

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
go build -o target/%s.exe ./main.go ./routers.go
if ($LASTEXITCODE -ne 0) {
    exit $LASTEXITCODE
}

# run app
./target/%s.exe
`, projectName, projectName)
	} else {
		return fmt.Sprintf(`# build app
go build -o target/%s main.go routers.go

code=$?
if [ $code -ne 0 ]; then
    exit $code
fi

# run app
target/%s
`, projectName, projectName)
	}
}

func GetBuildScript(projectName string) string {
	if IsWindows() {
		return fmt.Sprintf(`# build app
go build -o target/%s.exe ./main.go ./routers.go
`, projectName)
	} else {
		return fmt.Sprintf(`# build app
go build -o target/%s main.go routers.go
`, projectName)
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
