package models

import (
	"context"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/SpaceSlow/gophkeeper/generated/openapi"
	"github.com/SpaceSlow/gophkeeper/internal/domain/sensitive_records"
)

type SensitiveRecordModel struct {
	ctx                 context.Context
	client              *openapi.ClientWithResponses
	index               int
	sensitiveRecord     *sensitive_records.SensitiveRecord
	sensitiveRecordData *sensitive_records.SensitiveRecordData
}

func NewSensitiveRecordModel(
	ctx context.Context,
	client *openapi.ClientWithResponses,
	index int,
	sensitiveRecord *sensitive_records.SensitiveRecord,
) *SensitiveRecordModel {
	response, err := client.FetchSensitiveRecordWithIDWithResponse(ctx, sensitiveRecord.Id())
	if err != nil {
		return nil
	}

	data, _ := sensitive_records.NewSensitiveRecordData(sensitiveRecord.Id(), response.Body)

	return &SensitiveRecordModel{
		ctx:                 ctx,
		client:              client,
		index:               index,
		sensitiveRecord:     sensitiveRecord,
		sensitiveRecordData: data,
	}
}

func (m SensitiveRecordModel) Init() tea.Cmd {
	return nil
}

func (m SensitiveRecordModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			m := NewTableModel(m.ctx, m.client)
			return m, nil
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m SensitiveRecordModel) View() string {
	return fmt.Sprintf("Sensitive record #%d. %s\n\n%s", m.index, m.sensitiveRecord.Metadata(), m.sensitiveRecordData.Data())
}
