package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestPayload struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

func Ask(text string) string {
	keyFile, err := os.ReadFile("key.txt")
	if err != nil {
		fmt.Println("failed to read key from key.txt :/")
		fmt.Println(err)
		os.Exit(1)
	}

	url := "https://api.openai.com/v1/chat/completions"
	key := string(keyFile)

	payload := RequestPayload{
		Model:    "gpt-4",
		Messages: []Message{{Role: "user", Content: text}},
	}

	payloadText, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("json creation failed :/")
		fmt.Println(err)
		os.Exit(2)
	}

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer((payloadText)))
	if err != nil {
		fmt.Println("request creation failed :/")
		fmt.Println(err)
		os.Exit(3)
	}

	request.Header.Set("Authorization", "Bearer "+key)
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("request proccesing failed :/")
		fmt.Println(err)
		os.Exit(4)
	}
	defer response.Body.Close()

	return Parse(response)
}
