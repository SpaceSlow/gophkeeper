package models

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/SpaceSlow/gophkeeper/generated/openapi"
	"github.com/SpaceSlow/gophkeeper/internal/application/models/keys"
	"github.com/SpaceSlow/gophkeeper/internal/domain/sensitive_records"
)

type TextFormModel struct {
	ctx    context.Context
	client *openapi.ClientWithResponses

	textArea        textarea.Model
	metadata        textinput.Model
	isFocusTextArea bool

	keys keys.TextFormKeyMap
	help help.Model
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

	helpModel := help.New()
	helpModel.ShowAll = true
	return TextFormModel{
		ctx:             ctx,
		client:          client,
		textArea:        textArea,
		metadata:        metadata,
		isFocusTextArea: true,
		keys:            keys.TextFormKeys,
		help:            helpModel,
	}
}

func (m TextFormModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m TextFormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Enter):
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
		case key.Matches(msg, m.keys.PrevInput) || key.Matches(msg, m.keys.NextInput):
			m.nextInput()
		case key.Matches(msg, m.keys.Back):
			return NewTableModel(m.ctx, m.client), nil
		case key.Matches(msg, m.keys.Quit):
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
	form := fmt.Sprintf(` %s
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

	helpView := m.help.View(m.keys)
	height := 20 - strings.Count(form, "\n") - strings.Count(helpView, "\n")
	return "\n" + form + strings.Repeat("\n", height) + helpView
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
