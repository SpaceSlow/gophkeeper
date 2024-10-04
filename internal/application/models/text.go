package models

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/SpaceSlow/gophkeeper/generated/openapi"
	"github.com/SpaceSlow/gophkeeper/internal/domain/sensitive_records"
)

type TextFormModel struct {
	ctx    context.Context
	client *openapi.ClientWithResponses

	textArea        textarea.Model
	metadata        textinput.Model
	isFocusTextArea bool
}

func NewTextFormModel(ctx context.Context, client *openapi.ClientWithResponses) tea.Model {
	textArea := textarea.New()
	textArea.Placeholder = "Some text..."
	textArea.Focus()

	metadata := textinput.New()
	metadata.Placeholder = "some metadata"
	metadata.CharLimit = 100
	metadata.Width = 50
	metadata.Prompt = ""

	return TextFormModel{
		ctx:             ctx,
		client:          client,
		textArea:        textArea,
		metadata:        metadata,
		isFocusTextArea: true,
	}
}

func (m TextFormModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m TextFormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.isFocusTextArea {
				break
			}
			response, _ := m.client.PostSensitiveRecordWithResponse(m.ctx, openapi.PostSensitiveRecordJSONRequestBody{
				Metadata: m.metadata.Value(),
				Type:     openapi.Text,
			})
			var data bytes.Buffer
			enc := gob.NewEncoder(&data)

			text := sensitive_records.Text{
				Data: m.textArea.Value(),
			}
			enc.Encode(text)
			_, _ = m.client.PostSensitiveRecordDataWithBodyWithResponse(
				m.ctx,
				response.JSON201.Id,
				"application/octet-stream",
				&data,
			)
			return NewTextModel(m.ctx, m.client, &text, m.metadata.Value()), nil
		case tea.KeyShiftTab, tea.KeyTab:
			m.nextInput()
		case tea.KeyEsc:
			return NewTableModel(m.ctx, m.client), nil
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	}

	m.textArea, cmd = m.textArea.Update(msg)
	cmds = append(cmds, cmd)
	m.metadata, cmd = m.metadata.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m TextFormModel) View() string {
	return fmt.Sprintf(` %s
%s

 %s: %s

 %s
`,
		"Text",
		m.textArea.View(),
		"Metadata",
		m.metadata.View(),
		"Continue ->",
	)
}

func (m *TextFormModel) nextInput() {
	m.isFocusTextArea = !m.isFocusTextArea
	if m.isFocusTextArea {
		m.metadata.Blur()
		m.textArea.Focus()
		return
	}
	m.textArea.Blur()
	m.metadata.Focus()
}

type TextModel struct {
	ctx    context.Context
	client *openapi.ClientWithResponses

	text     *sensitive_records.Text
	metadata string
}

func NewTextModel(
	ctx context.Context,
	client *openapi.ClientWithResponses,
	text *sensitive_records.Text,
	metadata string,
) tea.Model {
	return TextModel{
		ctx:      ctx,
		client:   client,
		text:     text,
		metadata: metadata,
	}
}

func (m TextModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m TextModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m TextModel) View() string {
	return fmt.Sprintf(` %s
 %s

 %s: %s
`,
		"Text",
		m.text.Data,
		"Metadata",
		m.metadata,
	)
}
