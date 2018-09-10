package clients

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

var (
	webhook = os.Getenv("SlackWebhookURL")

	Slack = newSlackClient()
)

type slack struct {
	client *http.Client
}

func newSlackClient() *slack {
	return &slack{
		client: &http.Client{},
	}
}

// Send sends message to slack via webhook
func (s *slack) Send(message string) error {
	var jsonStr = []byte(fmt.Sprintf(`{"text": "%s"}`, message))

	req, err := http.NewRequest("POST", webhook, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}

	req.Header.Set("Content-type", "application/json")

	_, err = s.client.Do(req)
	return err
}
