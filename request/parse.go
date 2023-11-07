package request

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Parse(response *http.Response) string {
	textJSON, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("failed to read response body :/")
		fmt.Println(err)
		os.Exit(1)
	}

	var parsed ResponsePayload
	err = json.Unmarshal(textJSON, &parsed)
	if err != nil {
		fmt.Println("failed to decode response json :/")
		fmt.Println(err)
		fmt.Println("recieved json:\n" + string(textJSON))
		os.Exit(1)
	}

	return parsed.Choices[0].Message.Content
}
