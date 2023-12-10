package utils

import (
	"fmt"
	"os"
)

func Input(hint string) string {
	var value string

	fmt.Print(hint)
	_, err := fmt.Scan(&value)
	if err != nil {
		panic(err)
	}

	return value
}

func Save(path string, data string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		panic(err)
	}
}
