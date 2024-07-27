package llm

import (
	"context"
	"fmt"

	"github.com/henomis/lingoose/llm/antropic"
	"github.com/henomis/lingoose/thread"
	"github.com/isaacphi/codeassistantprogram/internal/models"
)

type LLM struct {
	model       string
	providerLLM *antropic.Antropic
	thread      *thread.Thread
}

func New(model string) *LLM {
	providerLLM := antropic.New().WithModel(model).WithStream(
		func(response string) {
			if response != antropic.EOS {
				fmt.Print(response)
			} else {
				fmt.Println()
			}
		},
	)

	t := thread.New()

	return &LLM{
		model:       model,
		providerLLM: providerLLM,
		thread:      t,
	}
}

func (llm *LLM) AddMessage(messageContents string, messageType string) error {
	switch messageType {
	case "user":
		llm.thread.AddMessage(
			thread.NewUserMessage().AddContent(
				thread.NewTextContent(messageContents),
			),
		)
		return nil
	case "assistant":
		llm.thread.AddMessage(
			thread.NewAssistantMessage().AddContent(
				thread.NewTextContent(messageContents),
			),
		)
		return nil
	default:
		return fmt.Errorf("messageType %v must be \"user\" or \"assistant\"", messageType)
	}
}

func (llm *LLM) GenerateResponse() (string, error) {
	err := llm.providerLLM.Generate(context.Background(), llm.thread)
	if err != nil {
		return "", fmt.Errorf("failed to generate response: %w", err)
	}
	return llm.thread.LastMessage().Contents[0].AsString(), nil
}

func (llm *LLM) LoadThread(t *models.Thread) error {
	llm.thread.ClearMessages()
	for _, messageID := range t.MessageIDs {
		message, err := models.LoadMessage(messageID)
		if err != nil {
			return fmt.Errorf("failed to load message %v: %w", messageID, err)
		}
		err = llm.AddMessage(message.Content, message.Type)
		if err != nil {
			return fmt.Errorf("failed to add message to LLM %v: %w", messageID, err)
		}
	}
	return nil
}
