package models

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/SpaceSlow/gophkeeper/internal/application/models/keys"
	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/SpaceSlow/gophkeeper/generated/openapi"
	"github.com/SpaceSlow/gophkeeper/internal/domain/sensitive_records"
)

type BinaryFormModel struct {
	ctx    context.Context
	client *openapi.ClientWithResponses

	filepicker    filepicker.Model
	metadataInput textinput.Model
	selectedFile  string

	err error
}

func NewBinaryFormModel(ctx context.Context, client *openapi.ClientWithResponses) tea.Model {
	fp := filepicker.New()
	fp.CurrentDirectory, _ = os.UserHomeDir()
	fp.Height = 15
	fp.ShowPermissions = false
	fp.AutoHeight = true

	metadataInput := textinput.New()
	metadataInput.CharLimit = 100
	metadataInput.Width = 50
	metadataInput.Placeholder = "some metadata"
	metadataInput.Prompt = ""

	m := BinaryFormModel{
		ctx:           ctx,
		client:        client,
		filepicker:    fp,
		metadataInput: metadataInput,
	}
	return m
}

type clearErrorMsg struct{}

func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return clearErrorMsg{}
	})
}

func (m BinaryFormModel) Init() tea.Cmd {
	return m.filepicker.Init()
}

func (m BinaryFormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.selectedFile == "" {
				break
			}
			fileData, _ := os.ReadFile(m.selectedFile)
			response, _ := m.client.PostSensitiveRecordWithResponse(m.ctx, openapi.PostSensitiveRecordJSONRequestBody{
				Metadata: m.metadataInput.Value(),
				Type:     openapi.Binary,
			})
			var data bytes.Buffer
			enc := gob.NewEncoder(&data)

			binary := sensitive_records.Binary{
				Data: fileData,
			}
			enc.Encode(binary)
			_, _ = m.client.PostSensitiveRecordDataWithBodyWithResponse(
				m.ctx,
				response.JSON201.Id,
				"application/octet-stream",
				&data,
			)
			return NewBinaryModel(m.ctx, m.client, &binary, m.metadataInput.Value()), nil
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEsc:
			return NewTableModel(m.ctx, m.client), nil
		}
	case clearErrorMsg:
		m.err = nil
	}

	var cmds []tea.Cmd
	var cmd tea.Cmd
	m.filepicker, cmd = m.filepicker.Update(msg)
	cmds = append(cmds, cmd)

	if didSelect, path := m.filepicker.DidSelectFile(msg); didSelect {
		if isDisable, _ := m.filepicker.DidSelectDisabledFile(msg); isDisable {
			m.err = errors.New(path + " is not valid.")
			return m, tea.Batch(cmd, clearErrorAfter(2*time.Second))
		}
		m.selectedFile = path
		m.metadataInput.Focus()
		return m, nil
	}

	m.metadataInput, cmd = m.metadataInput.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m BinaryFormModel) View() string {
	var s strings.Builder
	s.WriteString("\n  ")
	if m.err != nil {
		s.WriteString(m.filepicker.Styles.DisabledFile.Render(m.err.Error()))
	} else if m.selectedFile == "" {
		s.WriteString("Pick a file:")
		s.WriteString("\n\n" + m.filepicker.View() + "\n")
	} else {
		s.WriteString(fmt.Sprintf("%s: %s\n\n", "Selected file: ", m.filepicker.Styles.Selected.Render(m.selectedFile)))
		s.WriteString(fmt.Sprintf(" Metadata: %s\n", m.metadataInput.View()))
		s.WriteString(" NOTE: save filename in metadata\n\n")
		s.WriteString(" Continue ->")
	}
	return s.String()
}

type BinaryModel struct {
	ctx    context.Context
	client *openapi.ClientWithResponses

	binary    *sensitive_records.Binary
	metadata  string
	isSaved   bool
	pathInput textinput.Model

	keys keys.BinaryKeyMap
	help help.Model
}

func NewBinaryModel(ctx context.Context, client *openapi.ClientWithResponses, binary *sensitive_records.Binary, metadata string) tea.Model {
	pathInput := textinput.New()
	pathInput.CharLimit = 100
	pathInput.Width = 50
	pathInput.Placeholder = "file path"
	pathInput.Prompt = ""

	helpModel := help.New()
	helpModel.ShowAll = true
	return &BinaryModel{
		ctx:       ctx,
		client:    client,
		binary:    binary,
		metadata:  metadata,
		pathInput: pathInput,

		keys: keys.BinaryKeys,
		help: helpModel,
	}
}

func (m BinaryModel) Init() tea.Cmd {
	return nil
}

func (m BinaryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Enter):
			if !m.pathInput.Focused() {
				break
			}
			err := os.WriteFile(m.pathInput.Value(), m.binary.Data, 0664)
			if err != nil {
				break
			}
			m.isSaved = true
		case key.Matches(msg, m.keys.Save):
			m.pathInput.Focus()
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Back):
			return NewTableModel(m.ctx, m.client), nil
		}
	}
	var cmd tea.Cmd
	m.pathInput, cmd = m.pathInput.Update(msg)

	return m, cmd
}

func (m BinaryModel) View() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintf(" Metadata: %s\n", m.metadata))
	if m.pathInput.Focused() {
		s.WriteString("\n Path: ")
		s.WriteString(m.pathInput.View())
		s.WriteRune('\n')
	}
	if m.isSaved {
		s.WriteString(" File saved!")
	}
	binaryView := s.String()
	helpView := m.help.View(m.keys)
	height := 20 - strings.Count(binaryView, "\n") - strings.Count(helpView, "\n")
	return "\n" + binaryView + strings.Repeat("\n", height) + helpView
}
