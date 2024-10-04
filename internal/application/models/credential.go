package models

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/SpaceSlow/gophkeeper/generated/openapi"
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

	return CredentialFormModel{
		ctx:     ctx,
		client:  client,
		inputs:  inputs,
		focused: 0,
	}
}

func (m CredentialFormModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m CredentialFormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
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
		case tea.KeyShiftTab:
			m.prevInput()
		case tea.KeyTab:
			m.nextInput()
		case tea.KeyEsc:
			return NewTableModel(m.ctx, m.client), nil
		case tea.KeyCtrlC:
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
	return fmt.Sprintf(` %s

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
	}
}

func (m CredentialModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m CredentialModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			return NewTableModel(m.ctx, m.client), nil
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m CredentialModel) View() string {
	return fmt.Sprintf(` %s

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
}
