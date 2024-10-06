package models

import (
	"bytes"
	"context"
	"encoding/gob"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/SpaceSlow/gophkeeper/generated/openapi"
	"github.com/SpaceSlow/gophkeeper/internal/application/models/keys"
	"github.com/SpaceSlow/gophkeeper/internal/domain/sensitive_records"
)

type TableModel struct {
	ctx    context.Context
	client *openapi.ClientWithResponses
	table  table.Model

	sensitiveRecords []sensitive_records.SensitiveRecord

	keys keys.TableKeyMap
	help help.Model
}

func NewTableModel(
	ctx context.Context,
	client *openapi.ClientWithResponses,
) tea.Model {
	helpModel := help.New()
	helpModel.ShowAll = true
	model := &TableModel{
		ctx:    ctx,
		client: client,

		keys: keys.TableKeys,
		help: helpModel,
	}
	model.fetchSensitiveRecords()
	model.fillTable()
	return model
}

func (m TableModel) Init() tea.Cmd {
	return nil
}

func (m TableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Open):
			i, _ := strconv.Atoi(m.table.SelectedRow()[0])
			sensitiveRecord := m.sensitiveRecords[i-1]
			response, _ := m.client.FetchSensitiveRecordWithIDWithResponse(m.ctx, sensitiveRecord.Id())
			data := bytes.NewBuffer(response.Body)
			dec := gob.NewDecoder(data)
			switch openapi.SensitiveRecordTypeEnum(sensitiveRecord.Type()) {
			case openapi.PaymentCard:
				var paymentCard sensitive_records.PaymentCard
				dec.Decode(&paymentCard)
				return NewPaymentCardModel(m.ctx, m.client, &paymentCard, sensitiveRecord.Metadata()), nil
			case openapi.Text:
				var text sensitive_records.Text
				dec.Decode(&text)
				return NewTextModel(m.ctx, m.client, &text, sensitiveRecord.Metadata()), nil
			case openapi.Credential:
				var credential sensitive_records.Credential
				dec.Decode(&credential)
				return NewCredentialModel(m.ctx, m.client, &credential, sensitiveRecord.Metadata()), nil
			case openapi.Binary:
				var binary sensitive_records.Binary
				dec.Decode(&binary)
				return NewBinaryModel(m.ctx, m.client, &binary, sensitiveRecord.Metadata()), nil
			}
		case key.Matches(msg, m.keys.Create):
			return NewChoiceCreateSensitiveRecordModel(m.ctx, m.client), nil
		case key.Matches(msg, m.keys.Back):
			return NewMainModel(m.ctx, "https://localhost/api/"), nil
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m TableModel) View() string {
	table := m.table.View()

	helpView := m.help.View(m.keys)
	height := 20 - strings.Count(table, "\n") - strings.Count(helpView, "\n")
	return "\n" + table + strings.Repeat("\n", height) + helpView
}

func (m *TableModel) fetchSensitiveRecords() {
	records, _ := m.client.ListSensitiveRecordsWithResponse(m.ctx)

	sensitiveRecords := make([]sensitive_records.SensitiveRecord, 0, len(records.JSON200.SensitiveRecords))
	for _, record := range records.JSON200.SensitiveRecords {
		r, _ := sensitive_records.NewSensitiveRecord(record.Id, 0, string(record.Type), record.Metadata)
		sensitiveRecords = append(sensitiveRecords, *r)
	}

	m.sensitiveRecords = sensitiveRecords
}

func (m *TableModel) fillTable() {
	columns := []table.Column{
		{Title: "№", Width: 4},
		{Title: "metadata", Width: 70},
	}

	rows := make([]table.Row, 0, len(m.sensitiveRecords))

	for i, r := range m.sensitiveRecords {
		rows = append(rows, []string{strconv.Itoa(i + 1), r.Metadata()})
	}

	m.table = table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(15),
	)
}
