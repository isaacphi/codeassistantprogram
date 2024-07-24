package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func GetInput() (string, error) {
	model := NewTextInputModel()
	p := tea.NewProgram(model)

	finalModel, err := p.Run()
	if err != nil {
		return "", err
	}

	return finalModel.(TextInputModel).textinput.Value(), nil
}
