package models

import (
	"context"
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"

	"github.com/SpaceSlow/gophkeeper/generated/openapi"
)

type LoginModel struct {
	ctx     context.Context
	client  *openapi.ClientWithResponses
	address string

	inputs  []textinput.Model
	focused int
}

func NewLoginModel(ctx context.Context, client *openapi.ClientWithResponses, address string) tea.Model {
	var inputs = make([]textinput.Model, 2)
	inputs[username] = textinput.New()
	inputs[username].Placeholder = "username"
	inputs[username].Focus()
	inputs[username].CharLimit = 20
	inputs[username].Width = 20

	inputs[password] = textinput.New()
	inputs[password].Placeholder = "password"
	inputs[password].CharLimit = 20
	inputs[password].Width = 20

	return &LoginModel{
		ctx:     ctx,
		client:  client,
		address: address,

		inputs:  inputs,
		focused: 0,
	}
}

func (m *LoginModel) Init() tea.Cmd {
	return nil
}

func (m *LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyShiftTab, tea.KeyUp:
			m.prevInput()
		case tea.KeyTab, tea.KeyDown:
			m.nextInput()
		case tea.KeyEnter:
			m.authenticate(m.inputs[username].Value(), m.inputs[password].Value())
			return NewTableModel(m.ctx, m.client), nil
		case tea.KeyEsc:
			return NewMainModel(m.ctx, m.address), nil
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()

		for i := range m.inputs {
			m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
		}
		return m, tea.Batch(cmds...)
	}

	return m, nil
}

func (m *LoginModel) View() string {
	return fmt.Sprintf(
		`Login Form:
%s: %s
%s: %s
`,
		"username",
		m.inputs[username].View(),
		"password",
		m.inputs[password].View(),
	) + "\n"
}

func (m *LoginModel) authenticate(username, password string) {
	login, _ := m.client.PostLoginWithResponse(context.TODO(), openapi.PostLoginJSONRequestBody{
		Username: username,
		Password: password,
	})
	bearerAuth, _ := securityprovider.NewSecurityProviderBearerToken(login.JSON200.Token)
	m.client, _ = openapi.NewClientWithResponses(m.address, openapi.WithRequestEditorFn(bearerAuth.Intercept))
}

func (m *LoginModel) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

func (m *LoginModel) prevInput() {
	m.focused--
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}
