package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type TextInputModel struct {
	textinput textinput.Model
	err       error
}

func NewTextInputModel() TextInputModel {
	ti := textinput.New()
	ti.Placeholder = "Type something..."
	ti.Focus()

	return TextInputModel{
		textinput: ti,
		err:       nil,
	}
}

func (m TextInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m TextInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyEnter:
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.textinput.Width = msg.Width - 4
	}

	m.textinput, cmd = m.textinput.Update(msg)
	return m, cmd
}

func (m TextInputModel) View() string {
	return m.textinput.View()
}
