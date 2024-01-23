package utils

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"maps"
	"os"
	"path/filepath"
)

type Config struct {
	GsVersion string    `json:"gs-version"`
	App       AppConfig `json:"app"`
}

type AppConfig struct {
	Name         string              `json:"name"`
	Version      string              `json:"version"`
	CustomConfig bool                `json:"custom-config"`
	Main         string              `json:"main"`
	Target       string              `json:"target"`
	Scripts      map[string][]string `json:"scripts"`
}

var path2Md5 map[string]string

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

func NewConfig() Config {
	var config Config
	config.App.Scripts = make(map[string][]string)
	return config
}

func ReadConfig(path string) Config {
	jsonData, err := os.ReadFile(path)
	if err != nil {
		ThrowE(err)
	}
	config := NewConfig()
	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		ThrowE(err)
	}
	return config
}

func Mkdir(path string) {
	if err := os.Mkdir(path, fs.ModeDir); err != nil {
		ThrowE(err)
	}
}

func PrintE(err error) {
	fmt.Println("error:", err.Error())
}

func ThrowE(err error) {
	PrintE(err)
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

func CalcMd5(fpath string) (string, error) {
	file, err := os.Open(fpath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	md5Bytes := hash.Sum(nil)

	return hex.EncodeToString(md5Bytes), nil
}

func CheckModified() bool {
	neoPath2Md5 := make(map[string]string)

	if err := filepath.Walk(GetPath(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		name := info.Name()
		if info.IsDir() && name == ".git" {
			return filepath.SkipDir
		}
		if !info.IsDir() && name[max(0, len(name)-3):] == ".go" {
			if md5Str, err := CalcMd5(path); err != nil {
				ThrowE(err)
			} else {
				neoPath2Md5[path] = md5Str
			}
		}

		return nil
	}); err != nil {
		ThrowE(err)
	}

	result := !maps.Equal(neoPath2Md5, path2Md5)
	path2Md5 = neoPath2Md5
	return result
}
