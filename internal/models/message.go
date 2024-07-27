package models

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/isaacphi/codeassistantprogram/internal/config"
	"github.com/isaacphi/codeassistantprogram/internal/storage/fileio"
)

type Message struct {
	ID        string
	CreatedAt time.Time
	Content   string
	Type      string
}

func (m *Message) Save() error {
	basePath := config.DataDirectory
	return fileio.SaveYAML(m, basePath, "messages", m.ID)
}

func (m *Message) Delete() error {
	basePath := config.DataDirectory
	return fileio.DeleteFile(basePath, "messages", m.ID)
}

func NewMessage(content string, messageType string) (*Message, error) {
	if messageType != "user" && messageType != "assistant" {
		return nil, fmt.Errorf("messageType %v must be \"user\" or \"assistant\"", messageType)
	}

	id, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("failed to generate UUID: %w", err)
	}

	return &Message{
		ID:        id.String(),
		CreatedAt: time.Now(),
		Content:   content,
		Type:      messageType,
	}, nil
}

func LoadMessage(id string) (*Message, error) {
	basePath := config.DataDirectory
	var m Message
	err := fileio.LoadYAML(&m, basePath, "messages", id)
	if err != nil {
		return nil, fmt.Errorf("failed to load message: %w", err)
	}
	return &m, nil
}
