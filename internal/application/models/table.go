package models

import (
	"context"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/SpaceSlow/gophkeeper/generated/openapi"
	"github.com/SpaceSlow/gophkeeper/internal/domain/sensitive_records"
)

type TableModel struct {
	ctx    context.Context
	client *openapi.ClientWithResponses
	table  table.Model

	sensitiveRecords []sensitive_records.SensitiveRecord
}

func NewTableModel(
	ctx context.Context,
	client *openapi.ClientWithResponses,
) tea.Model {
	records, _ := client.ListSensitiveRecordsWithResponse(ctx)

	sensitiveRecords := make([]sensitive_records.SensitiveRecord, 0, len(records.JSON200.SensitiveRecords))
	for _, record := range records.JSON200.SensitiveRecords {
		r, _ := sensitive_records.NewSensitiveRecord(record.Id, 0, string(record.Type), record.Metadata)
		sensitiveRecords = append(sensitiveRecords, *r)
	}

	columns := []table.Column{
		{Title: "№", Width: 4},
		{Title: "metadata", Width: 70},
	}

	rows := make([]table.Row, 0, len(sensitiveRecords))

	for i, r := range sensitiveRecords {
		rows = append(rows, []string{strconv.Itoa(i + 1), r.Metadata()})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	return &TableModel{
		ctx:              ctx,
		client:           client,
		table:            t,
		sensitiveRecords: sensitiveRecords,
	}
}

func (m TableModel) Init() tea.Cmd {
	return nil
}

func (m TableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			i, _ := strconv.Atoi(m.table.SelectedRow()[0])
			return NewSensitiveRecordModel(m.ctx, m.client, i-1, &m.sensitiveRecords[i-1]), cmd
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m TableModel) View() string {
	return m.table.View() + "\n  " + m.table.HelpView() + "\n"
}

func (m TableModel) FetchSensitiveRecords() {
	records, _ := m.client.ListSensitiveRecordsWithResponse(m.ctx)

	sensitiveRecords := make([]sensitive_records.SensitiveRecord, 0, len(records.JSON200.SensitiveRecords))
	for _, record := range records.JSON200.SensitiveRecords {
		r, _ := sensitive_records.NewSensitiveRecord(record.Id, 0, string(record.Type), record.Metadata)
		sensitiveRecords = append(sensitiveRecords, *r)
	}

	columns := []table.Column{
		{Title: "№", Width: 4},
		{Title: "metadata", Width: 70},
	}

	rows := make([]table.Row, 0, len(sensitiveRecords))

	for i, r := range sensitiveRecords {
		rows = append(rows, []string{strconv.Itoa(i + 1), r.Metadata()})
	}

	m.table = table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)
}
