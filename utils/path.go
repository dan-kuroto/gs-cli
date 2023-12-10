package utils

import (
	"os"
	"path/filepath"
)

func GetPath(paths ...string) string {
	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	paths = append([]string{root}, paths...)
	return filepath.Join(paths...)
}
