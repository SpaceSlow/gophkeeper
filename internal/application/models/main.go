package models

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/SpaceSlow/gophkeeper/generated/openapi"
)

type MainModel struct {
	ctx     context.Context
	client  *openapi.ClientWithResponses
	address string
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
	}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "1":
			return NewRegisterModel(m.ctx, m.client, m.address), nil
		case "2":
			return NewLoginModel(m.ctx, m.client, m.address), nil
		case "ctrl+c", "esc":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m MainModel) View() string {
	return "Register - 1\nLogin - 2"
}
