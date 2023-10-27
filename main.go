package main

import (
	"askgpt/request"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("error, please ask a question")
		os.Exit(100)
	}

	fmt.Println(request.Ask(strings.Join(args[1:], " ")))
}
