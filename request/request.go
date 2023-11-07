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

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("failed to get home directory :/")
		fmt.Println(err)
		os.Exit(1)
	}

	conf, err := os.ReadFile(home + "/.local/share/askgpt/config.json")
	if err != nil {
		fmt.Println("could not locate config.json file in ~/.local/share/askgpt, please create it :/")
		fmt.Println(err)
		os.Exit(1)
	}

	config := Configuration{
		Model: "gpt-4-1106-preview",
	}
	err = json.Unmarshal(conf, &config)
	if err != nil {
		fmt.Println("bad config file :/")
		fmt.Println(err)
		os.Exit(1)
	}

	payload := RequestPayload{
		Model:    config.Model,
		Messages: []Message{{Role: "user", Content: text}},
	}

	payloadText, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("json creation failed :/")
		fmt.Println(err)
		os.Exit(1)
	}

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer((payloadText)))
	if err != nil {
		fmt.Println("request creation failed :/")
		fmt.Println(err)
		os.Exit(1)
	}

	request.Header.Set("Authorization", "Bearer "+key)
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("request proccesing failed :/")
		fmt.Println(err)
		os.Exit(1)
	}
	defer response.Body.Close()

	return Parse(response)
}
