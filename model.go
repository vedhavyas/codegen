package codegen

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
	"time"
)

type Model struct {
	Name    string
	Retries int
	client  *openai.Client
}

func NewModel(model string, retries int, accessToken string) Model {
	return Model{
		Name:    model,
		Retries: retries,
		client:  openai.NewClient(accessToken),
	}
}

func (m Model) RespondJSON(request openai.ChatCompletionRequest, response any) error {
	resp, err := m.respondWithRetry(request)
	if err != nil {
		return err
	}

	dec := json.NewDecoder(bytes.NewReader([]byte(resp.Choices[0].Message.Content)))
	err = dec.Decode(response)
	if err != nil {
		logrus.Errorf("Prompt: %+v", request)
		logrus.Errorf("Response: %+v", resp)
		return err
	}

	return nil
}

func (m Model) respondWithRetry(request openai.ChatCompletionRequest) (resp openai.ChatCompletionResponse, err error) {
	for i := 0; i < m.Retries; i++ {
		resp, err = m.client.CreateChatCompletion(context.Background(), request)
		if err == nil {
			return resp, nil
		}

		logrus.Debugln("Sleeping for a bit")
		time.Sleep(1 * time.Second)
		continue
	}

	return resp, fmt.Errorf("failed to get response")
}

func (m Model) Respond(request openai.ChatCompletionRequest) (string, error) {
	resp, err := m.respondWithRetry(request)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, err
}
