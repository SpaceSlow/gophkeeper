package models

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/SpaceSlow/gophkeeper/generated/openapi"
	"github.com/SpaceSlow/gophkeeper/internal/application/models/keys"
)

const (
	username = iota
	password
	repeatedPassword
)

type RegisterModel struct {
	ctx     context.Context
	client  *openapi.ClientWithResponses
	address string

	inputs  []textinput.Model
	focused int

	keys keys.RegisterKeyMap
	help help.Model
}

func NewRegisterModel(ctx context.Context, client *openapi.ClientWithResponses, address string) tea.Model {
	var inputs = make([]textinput.Model, 3)
	inputs[username] = textinput.New()
	inputs[username].Placeholder = "username"
	inputs[username].Focus()
	inputs[username].CharLimit = 20
	inputs[username].Width = 20

	inputs[password] = textinput.New()
	inputs[password].Placeholder = "password"
	inputs[password].CharLimit = 20
	inputs[password].Width = 20

	inputs[repeatedPassword] = textinput.New()
	inputs[repeatedPassword].Placeholder = "repeated password"
	inputs[repeatedPassword].CharLimit = 20
	inputs[repeatedPassword].Width = 20

	helpModel := help.New()
	helpModel.ShowAll = true
	return &RegisterModel{
		ctx:     ctx,
		client:  client,
		address: address,

		inputs:  inputs,
		focused: 0,

		keys: keys.RegisterKeys,
		help: helpModel,
	}
}

func (m *RegisterModel) Init() tea.Cmd {
	return nil
}

func (m *RegisterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if !m.register(m.inputs[username].Value(), m.inputs[password].Value(), m.inputs[repeatedPassword].Value()) {
				return m, nil
			}
			return NewLoginModel(m.ctx, m.client, m.address), nil
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

func (m *RegisterModel) View() string {
	registerForm := fmt.Sprintf(
		` Register Form:
 %s: %s
 %s: %s
 %s: %s
`,
		"username",
		m.inputs[username].View(),
		"password",
		m.inputs[password].View(),
		"repeat",
		m.inputs[repeatedPassword].View(),
	) + "\n"

	helpView := m.help.View(m.keys)
	height := 20 - strings.Count(registerForm, "\n") - strings.Count(helpView, "\n")
	return "\n" + registerForm + strings.Repeat("\n", height) + helpView
}

func (m *RegisterModel) register(username, password, repeatedPassword string) bool {
	register, _ := m.client.PostRegisterWithResponse(context.TODO(), openapi.PostRegisterJSONRequestBody{
		Username:         username,
		Password:         password,
		RepeatedPassword: repeatedPassword,
	})
	return register.Status() == "200"
}

func (m *RegisterModel) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

func (m *RegisterModel) prevInput() {
	m.focused--
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}
