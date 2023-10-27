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
		os.Exit(1)
	}

	keyFile, err := os.ReadFile("./key.txt")
	if err != nil {
		fmt.Println("failed to read key from key.txt :/")
		fmt.Println(err)
		os.Exit(2)
	}
	key := string(keyFile)

	fmt.Println(request.Ask(strings.Join(args[1:], " "), key))
}
