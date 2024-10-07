package models

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/SpaceSlow/gophkeeper/generated/openapi"
	"github.com/SpaceSlow/gophkeeper/internal/application/models/keys"
	"github.com/SpaceSlow/gophkeeper/internal/domain/sensitive_records"
)

const (
	credentialUsername = iota
	credentialPassword
	credentialMetadata
)

type CredentialFormModel struct {
	ctx    context.Context
	client *openapi.ClientWithResponses

	inputs  []textinput.Model
	focused int

	keys keys.CredentialFormKeyMap
	help help.Model
}

func NewCredentialFormModel(ctx context.Context, client *openapi.ClientWithResponses) tea.Model {
	var inputs = make([]textinput.Model, 3)
	inputs[credentialUsername] = textinput.New()
	inputs[credentialUsername].Placeholder = "Username"
	inputs[credentialUsername].Focus()
	inputs[credentialUsername].CharLimit = 50
	inputs[credentialUsername].Width = 20
	inputs[credentialUsername].Prompt = ""

	inputs[credentialPassword] = textinput.New()
	inputs[credentialPassword].Placeholder = "Password"
	inputs[credentialPassword].CharLimit = 50
	inputs[credentialPassword].Width = 20
	inputs[credentialPassword].Prompt = ""

	inputs[credentialMetadata] = textinput.New()
	inputs[credentialMetadata].Placeholder = "some metadata"
	inputs[credentialMetadata].CharLimit = 100
	inputs[credentialMetadata].Width = 50
	inputs[credentialMetadata].Prompt = ""

	helpModel := help.New()
	helpModel.ShowAll = true
	return CredentialFormModel{
		ctx:     ctx,
		client:  client,
		inputs:  inputs,
		focused: 0,

		keys: keys.CredentialFormKeys,
		help: helpModel,
	}
}

func (m CredentialFormModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m CredentialFormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Enter):
			if m.focused < len(m.inputs)-1 {
				m.nextInput()
				break
			}
			response, _ := m.client.PostSensitiveRecordWithResponse(m.ctx, openapi.PostSensitiveRecordJSONRequestBody{
				Metadata: m.inputs[credentialMetadata].Value(),
				Type:     openapi.Credential,
			})
			var data bytes.Buffer
			enc := gob.NewEncoder(&data)

			credential := sensitive_records.Credential{
				Username: m.inputs[credentialUsername].Value(),
				Password: m.inputs[credentialPassword].Value(),
			}
			enc.Encode(credential)
			_, _ = m.client.PostSensitiveRecordDataWithBodyWithResponse(
				m.ctx,
				response.JSON201.Id,
				"application/octet-stream",
				&data,
			)
			return NewCredentialModel(m.ctx, m.client, &credential, m.inputs[credentialMetadata].Value()), nil
		case key.Matches(msg, m.keys.PrevInput):
			m.prevInput()
		case key.Matches(msg, m.keys.NextInput):
			m.nextInput()
		case key.Matches(msg, m.keys.Back):
			return NewTableModel(m.ctx, m.client), nil
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m CredentialFormModel) View() string {
	form := fmt.Sprintf(` %s

 %s: %s
 %s: %s

 %s: %s

 %s
`,
		"Credential",
		"Username",
		m.inputs[credentialUsername].View(),
		"Password",
		m.inputs[credentialPassword].View(),
		"Metadata",
		m.inputs[credentialMetadata].View(),
		"Continue ->",
	)

	helpView := m.help.View(m.keys)
	height := 20 - strings.Count(form, "\n") - strings.Count(helpView, "\n")
	return "\n" + form + strings.Repeat("\n", height) + helpView
}

func (m *CredentialFormModel) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

func (m *CredentialFormModel) prevInput() {
	m.focused--
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}

type CredentialModel struct {
	ctx    context.Context
	client *openapi.ClientWithResponses

	credential *sensitive_records.Credential
	metadata   string

	keys keys.BackQuitKeyMap
	help help.Model
}

func NewCredentialModel(
	ctx context.Context,
	client *openapi.ClientWithResponses,
	credential *sensitive_records.Credential,
	metadata string,
) tea.Model {
	return CredentialModel{
		ctx:        ctx,
		client:     client,
		credential: credential,
		metadata:   metadata,
		keys:       keys.BackQuitKeys,
		help:       help.New(),
	}
}

func (m CredentialModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m CredentialModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Back):
			return NewTableModel(m.ctx, m.client), nil
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m CredentialModel) View() string {
	form := fmt.Sprintf(` %s

 %s: %s
 %s: %s

 %s: %s
`,
		"Credential",
		"Username",
		m.credential.Username,
		"Password",
		m.credential.Password,
		"Metadata",
		m.metadata,
	)

	helpView := m.help.View(m.keys)
	height := 20 - strings.Count(form, "\n") - strings.Count(helpView, "\n")
	return "\n" + form + strings.Repeat("\n", height) + helpView
}
