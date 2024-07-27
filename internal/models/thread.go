package models

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/isaacphi/codeassistantprogram/internal/config"
	fileio "github.com/isaacphi/codeassistantprogram/internal/storage/fileio"
)

type Thread struct {
	ID         string
	CreatedAt  time.Time
	Name       string
	MessageIDs []string
}

func (t *Thread) Save() error {
	basePath := config.DataDirectory
	name := t.ID
	if t.Name != "" {
		name = t.Name
	}
	return fileio.SaveYAML(t, basePath, "threads", name)
}

func (t *Thread) Delete() error {
	basePath := config.DataDirectory
	err := fileio.DeleteFile(basePath, "threads", t.ID)
	if err != nil {
		err = fileio.DeleteFile(basePath, "threads", t.Name)
		if err != nil {
			return fmt.Errorf("failed to delete thread %q: %w", t.ID, err)
		}
	}
	return nil
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
		message, err := LoadMessage(messageID)
		if err != nil {
			return err
		}
		fmt.Printf("%v: %v\n", message.Type, message.Content)
	}
	return nil
}

func (t *Thread) SetCurrent() error {
	err := fileio.SaveFile(config.DataDirectory, "HEAD", t.GetName())
	if err != nil {
		return fmt.Errorf("failed to save thread %q: %w", t.GetName(), err)
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

func LoadThread(id string) (*Thread, error) {
	basePath := config.DataDirectory
	var t Thread
	err := fileio.LoadYAML(&t, basePath, "threads", id)
	return &t, err
}

func ListThreads() ([]string, error) {
	basePath := config.DataDirectory
	return fileio.ListFiles(basePath, "threads")
}

func GetCurrentThread() (*Thread, error) {
	threadName, err := fileio.LoadFile(config.DataDirectory, "HEAD")
	if err != nil {
		return nil, err
	}
	thread, err := LoadThread(threadName)
	if err != nil {
		return nil, err
	}
	return thread, nil
}
