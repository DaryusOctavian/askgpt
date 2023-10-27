package request

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ResponseChoice struct {
	Index        uint64  `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type TokenUsage struct {
	PromptTokens     uint64 `json:"prompt_tokens"`
	CompletionTokens uint64 `json:"completion_tokens"`
	TotalTokens      uint64 `json:"total_tokens"`
}

type ResponsePayload struct {
	ID      string           `json:"id"`
	Object  string           `json:"object"`
	Created uint64           `json:"created"`
	Model   string           `json:"model"`
	Choices []ResponseChoice `json:"choices"`
	Usage   TokenUsage       `json:"usage"`
}

func Parse(response *http.Response) string {
	textJSON, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("failed to read response body :/")
		fmt.Println(err)
		os.Exit(6)
	}

	var parsed ResponsePayload
	err = json.Unmarshal(textJSON, &parsed)
	if err != nil {
		fmt.Println("failed to decode response json :/")
		fmt.Println(err)
		fmt.Println("recieved json:\n" + string(textJSON))
		os.Exit(7)
	}

	return parsed.Choices[0].Message.Content
}
