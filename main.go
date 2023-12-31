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

	key := os.Getenv("OPENAI_API_KEY")

	fmt.Println(request.Ask(strings.Join(args[1:], " "), key))
}
