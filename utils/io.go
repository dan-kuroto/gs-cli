package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Config struct {
	GsVersion string    `json:"gs-version"`
	App       AppConfig `json:"app"`
}

type AppConfig struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	CustomConfig bool   `json:"custom-config"`
	Main         string `json:"main"`
	Target       string `json:"target"`
}

func Input(hint string) string {
	var value string

	fmt.Print(hint)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		value = scanner.Text()
	} else {
		ThrowE(scanner.Err())
	}

	return value
}

func Save(path string, data string) {
	file, err := os.Create(path)
	if err != nil {
		ThrowE(err)
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		ThrowE(err)
	}
}

func Read(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		ThrowE(err)
	}
	return string(data)
}

func ReadConfig(path string) Config {
	jsonData, err := os.ReadFile(path)
	if err != nil {
		ThrowE(err)
	}
	var config Config
	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		ThrowE(err)
	}
	return config
}

func AddPackageToMainGo(projectName, mainPath, packageName string) {
	lines := make([]string, 0, 8)
	added := false
	replacer := strings.NewReplacer(
		"\t", "",
		"\n", "",
		"\v", "",
		"\f", "",
		"\r", "",
		" ", "",
		"\x85", "",
		"\xA0", "",
	)
	for _, line := range strings.Split(Read(mainPath), "\n") {
		lines = append(lines, line)
		if !added && replacer.Replace(line) == "packagemain" {
			added = true
			lines = append(lines, "", fmt.Sprintf(`import _ "%s/%s"`, projectName, packageName))
		}
	}
	Save(mainPath, strings.Join(lines, "\n"))
}

func Mkdir(path string) {
	if err := os.Mkdir(path, fs.ModeDir); err != nil {
		ThrowE(err)
	}
}

func ThrowE(err error) {
	fmt.Println("error:", err.Error())
	os.Exit(1)
}

func ThrowS(msg string, args ...any) {
	ThrowE(fmt.Errorf(msg, args...))
}

func AssertNotEmpty(name string, value any) {
	switch value := value.(type) {
	case string:
		if value == "" {
			ThrowS("%s can not be empty", name)
		}
	case []any:
		if len(value) == 0 {
			ThrowS("%s can not be empty", name)
		}
	default:
		ThrowS("Unsupported type %T", value)
	}
}

func CheckModified() bool {
	fileInfo, err := os.Stat("main.go")
	if err != nil {
		ThrowE(err)
	}
	modTime := fileInfo.ModTime()
	fmt.Printf("fileInfo.ModTime(): %v\n", modTime)
	fmt.Printf("modTime.Format(time.RFC3339): %v\n", modTime.Format(time.RFC3339))
	fmt.Printf("modTime.UnixMilli(): %v\n", modTime.UnixMilli())

	if err := filepath.Walk(GetPath(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if info.Name() == ".git" {
				return filepath.SkipDir
			}
			fmt.Println("dir:", path)
		} else {
			fmt.Println("file:", path)
		}

		return nil
	}); err != nil {
		ThrowE(err)
	}

	return false
}
