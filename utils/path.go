package utils

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetPath(paths ...string) string {
	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	paths = append([]string{root}, paths...)
	return filepath.Join(paths...)
}

func IsWindows() bool {
	return runtime.GOOS == "windows"
}

func GetScriptName(name string) string {
	if IsWindows() {
		return name + ".ps1"
	} else {
		return name + ".sh"
	}
}
