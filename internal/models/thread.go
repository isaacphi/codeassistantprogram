package models

import (
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/isaacphi/codeassistantprogram/internal/config"
	"github.com/isaacphi/codeassistantprogram/internal/storage/fileio"
	"time"
)

type Thread struct {
	ID        string
	CreatedAt time.Time
	Name      string
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

func NewThread(name string) (*Thread, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("Failed to generate UUID: %w", err)
	}

	return &Thread{
		ID:        id.String(),
		CreatedAt: time.Now(),
		Name:      name,
	}, nil
}

func LoadThread(id string, basePath string) (*Thread, error) {
	var thread Thread
	err := fileio.LoadYAML(&thread, basePath, "threads", id)
	return &thread, err
}

func ListThreads(basePath string) ([]string, error) {
	return fileio.ListFiles(basePath, "threads")
}

func SetCurrentThread(nameOrId string) (*Thread, error) {
	thread, err := LoadThread(nameOrId, config.DataDirectory)
	if err != nil {
		return nil, fmt.Errorf("Couldn't find thread %v\n%v", nameOrId, err)
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
