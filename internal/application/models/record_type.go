package models

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/SpaceSlow/gophkeeper/generated/openapi"
)

const listHeight = 14

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := func(s ...string) string {
		return "  " + strings.Join(s, " ")
	}
	if index == m.Index() {
		fn = func(s ...string) string {
			return "> " + strings.Join(s, " ")
		}
	}

	fmt.Fprint(w, fn(str))
}

type ChoiceCreateSensitiveRecordModel struct {
	ctx    context.Context
	client *openapi.ClientWithResponses

	list list.Model
}

func (m ChoiceCreateSensitiveRecordModel) Init() tea.Cmd {
	return nil
}

func (m ChoiceCreateSensitiveRecordModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			switch openapi.SensitiveRecordTypeEnum(m.list.SelectedItem().(item)) {
			case openapi.PaymentCard:
				return NewPaymentCardFormModel(m.ctx, m.client), nil
			}
		case "esc":
			return NewTableModel(m.ctx, m.client), nil
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ChoiceCreateSensitiveRecordModel) View() string {
	return m.list.View()
}

func NewChoiceCreateSensitiveRecordModel(ctx context.Context, client *openapi.ClientWithResponses) ChoiceCreateSensitiveRecordModel {
	items := []list.Item{
		item(openapi.PaymentCard),
		item(openapi.Credential),
		item(openapi.Text),
		item(openapi.Binary),
	}

	const defaultWidth = 20

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "What do you want to create sensitive record?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	return ChoiceCreateSensitiveRecordModel{ctx: ctx, client: client, list: l}
}
