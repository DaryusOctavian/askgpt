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

type Payload struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

func Ask() string {
	keyFile, err := os.ReadFile("key.txt")
	if err != nil {
		fmt.Println("failed to read key from key.txt :/")
		os.Exit(1)
	}

	url := "https://api.openai.com/v1/chat/completions"
	key := string(keyFile)

	payload := Payload{
		Model:    "gpt-4",
		Messages: []Message{{Role: "user", Content: "can you say hello?"}},
	}

	payloadText, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("json creation failed :/")
		os.Exit(2)
	}

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer((payloadText)))
	if err != nil {
		fmt.Println("request creation failed :/")
		os.Exit(3)
	}

	request.Header.Set("Authorization", "Bearer "+key)
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("request proccesing failed :/")
		os.Exit(4)
	}
	defer response.Body.Close()

	return Parse(request)
}
