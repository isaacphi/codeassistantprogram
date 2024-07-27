package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func GetInput() (string, error) {
	model := NewTextInputModel()
	p := tea.NewProgram(model)

	finalModel, err := p.Run()
	if err != nil {
		return "", fmt.Errorf("failed to get input: %w", err)
	}

	return finalModel.(TextInputModel).textinput.Value(), nil
}
