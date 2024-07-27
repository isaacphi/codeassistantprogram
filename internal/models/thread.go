package models

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/isaacphi/codeassistantprogram/internal/config"
	"github.com/isaacphi/codeassistantprogram/internal/storage/fileio"
)

type Thread struct {
	ID         string
	CreatedAt  time.Time
	Name       string
	MessageIDs []string
}

func (t *Thread) Save(basePath string) error {
	name := t.ID
	if t.Name != "" {
		name = t.Name
	}
	return fileio.SaveYAML(t, basePath, "threads", name)
}

func (t *Thread) Delete(basePath string) error {
	return fileio.DeleteFile(basePath, "threads", t.ID)
}

func (t *Thread) GetName() string {
	if t.Name != "" {
		return t.Name
	}
	return t.ID
}

func (t *Thread) AddMessage(message *Message) {
	t.MessageIDs = append(t.MessageIDs, message.ID)
}

func (t *Thread) View() error {
	for _, messageID := range t.MessageIDs {
		message, err := LoadMessage(messageID, config.DataDirectory)
		if err != nil {
			return err
		}
		fmt.Printf("%v: %v\n", message.Type, message.Content)
	}
	return nil
}

func NewThread(name string) (*Thread, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("failed to generate UUID: %w", err)
	}

	return &Thread{
		ID:        id.String(),
		CreatedAt: time.Now(),
		Name:      name,
	}, nil
}

func LoadThread(id string, basePath string) (*Thread, error) {
	var t Thread
	err := fileio.LoadYAML(&t, basePath, "threads", id)
	return &t, err
}

func ListThreads(basePath string) ([]string, error) {
	return fileio.ListFiles(basePath, "threads")
}

func SetCurrentThread(nameOrId string) (*Thread, error) {
	// TODO: pass in Thread, not name
	thread, err := LoadThread(nameOrId, config.DataDirectory)
	if err != nil {
		return nil, fmt.Errorf("failed to load thread %q: %w", nameOrId, err)
	}
	fileio.SaveFile(config.DataDirectory, "HEAD", thread.GetName())
	return thread, nil
}

func GetCurrentThread() (*Thread, error) {
	threadName, err := fileio.LoadFile(config.DataDirectory, "HEAD")
	if err != nil {
		return nil, err
	}
	thread, err := LoadThread(threadName, config.DataDirectory)
	if err != nil {
		return nil, err
	}
	return thread, nil
}
