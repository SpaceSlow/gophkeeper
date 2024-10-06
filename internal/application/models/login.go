package models

import (
	"context"
	"fmt"
	"github.com/SpaceSlow/gophkeeper/internal/application/models/keys"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"strings"

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

	keys keys.LoginKeyMap
	help help.Model
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

	helpModel := help.New()
	helpModel.ShowAll = true
	return &LoginModel{
		ctx:     ctx,
		client:  client,
		address: address,

		inputs:  inputs,
		focused: 0,

		keys: keys.LoginKeys,
		help: helpModel,
	}
}

func (m *LoginModel) Init() tea.Cmd {
	return nil
}

func (m *LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.PrevInput):
			m.prevInput()
		case key.Matches(msg, m.keys.NextInput):
			m.nextInput()
		case key.Matches(msg, m.keys.Enter):
			if m.focused < len(m.inputs)-1 {
				m.nextInput()
				break
			}
			m.authenticate(m.inputs[username].Value(), m.inputs[password].Value())
			return NewTableModel(m.ctx, m.client), nil
		case key.Matches(msg, m.keys.Back):
			return NewMainModel(m.ctx, m.address), nil
		case key.Matches(msg, m.keys.Quit):
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
	loginForm := fmt.Sprintf(
		` Login Form:
 %s: %s
 %s: %s
`,
		"username",
		m.inputs[username].View(),
		"password",
		m.inputs[password].View(),
	) + "\n"

	helpView := m.help.View(m.keys)
	height := 20 - strings.Count(loginForm, "\n") - strings.Count(helpView, "\n")
	return "\n" + loginForm + strings.Repeat("\n", height) + helpView
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
