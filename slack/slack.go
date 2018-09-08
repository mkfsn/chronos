package slack

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

var (
	webhook = os.Getenv("SlackWebhookURL")
)

func Send(message string) error {
	var jsonStr = []byte(fmt.Sprintf(`{"text": "%s"}`, message))

	client := &http.Client{}

	req, err := http.NewRequest("POST", webhook, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}

	req.Header.Set("Content-type", "application/json")

	_, err = client.Do(req)
	return err
}
