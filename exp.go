package main

import (
	"fmt"

	"github.com/dan-kuroto/gs-cli/utils"
)

func main() {
	if utils.CheckModified() {
		fmt.Println("modified")
	} else {
		fmt.Println("not modified")
	}
}
