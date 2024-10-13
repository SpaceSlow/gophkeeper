package models

import (
	"context"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/SpaceSlow/gophkeeper/generated/openapi"
	"github.com/SpaceSlow/gophkeeper/internal/application/models/keys"
)

type MainModel struct {
	ctx     context.Context
	client  *openapi.ClientWithResponses
	address string

	keys keys.MainKeyMap
	help help.Model
}

func NewMainModel(ctx context.Context, address string) tea.Model {
	client, err := openapi.NewClientWithResponses(address)
	if err != nil {
		return nil
	}

	return MainModel{
		ctx:     ctx,
		client:  client,
		address: address,

		keys: keys.MainKeys,
		help: help.New(),
	}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Register):
			return NewRegisterModel(m.ctx, m.client, m.address), nil
		case key.Matches(msg, m.keys.Login):
			return NewLoginModel(m.ctx, m.client, m.address), nil
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m MainModel) View() string {
	text := " 1. Register\n 2. Login"
	helpView := m.help.View(m.keys)
	height := 20 - strings.Count(text, "\n") - strings.Count(helpView, "\n")
	return "\n" + text + strings.Repeat("\n", height) + helpView
}
