package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func Ask(text string, key string) string {
	url := "https://api.openai.com/v1/chat/completions"

	payload := RequestPayload{
		Model:    "gpt-4-1106-preview",
		Messages: []Message{{Role: "user", Content: text}},
	}

	payloadText, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("json creation failed :/")
		fmt.Println(err)
		os.Exit(3)
	}

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer((payloadText)))
	if err != nil {
		fmt.Println("request creation failed :/")
		fmt.Println(err)
		os.Exit(4)
	}

	request.Header.Set("Authorization", "Bearer "+key)
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("request proccesing failed :/")
		fmt.Println(err)
		os.Exit(5)
	}
	defer response.Body.Close()

	return Parse(response)
}
