package models

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/SpaceSlow/gophkeeper/generated/openapi"
)

type (
	errMsg error
)

const (
	ccn = iota
	exp
	cvv
)

type PaymentCardModel struct {
	ctx    context.Context
	client *openapi.ClientWithResponses

	inputs  []textinput.Model
	focused int
	err     error
}

func ccnValidator(s string) error {
	if len(s) > 16+3 {
		return fmt.Errorf("CCN is too long")
	}

	if len(s) == 0 || len(s)%5 != 0 && (s[len(s)-1] < '0' || s[len(s)-1] > '9') {
		return fmt.Errorf("CCN is invalid")
	}

	if len(s)%5 == 0 && s[len(s)-1] != ' ' {
		return fmt.Errorf("CCN must separate groups with spaces")
	}

	c := strings.ReplaceAll(s, " ", "")
	_, err := strconv.ParseInt(c, 10, 64)

	return err
}

func expValidator(s string) error {
	e := strings.ReplaceAll(s, "/", "")
	_, err := strconv.ParseInt(e, 10, 64)
	if err != nil {
		return fmt.Errorf("EXP is invalid")
	}

	if len(s) >= 3 && (strings.Index(s, "/") != 2 || strings.LastIndex(s, "/") != 2) {
		return fmt.Errorf("EXP is invalid")
	}

	return nil
}

func cvvValidator(s string) error {
	_, err := strconv.ParseInt(s, 10, 64)
	return err
}

func NewPaymentCardModel(ctx context.Context, client *openapi.ClientWithResponses) tea.Model {
	var inputs = make([]textinput.Model, 3)
	inputs[ccn] = textinput.New()
	inputs[ccn].Placeholder = "4505 **** **** 1234"
	inputs[ccn].Focus()
	inputs[ccn].CharLimit = 20
	inputs[ccn].Width = 30
	inputs[ccn].Prompt = ""
	inputs[ccn].Validate = ccnValidator

	inputs[exp] = textinput.New()
	inputs[exp].Placeholder = "MM/YY "
	inputs[exp].CharLimit = 5
	inputs[exp].Width = 5
	inputs[exp].Prompt = ""
	inputs[exp].Validate = expValidator

	inputs[cvv] = textinput.New()
	inputs[cvv].Placeholder = "XXX"
	inputs[cvv].CharLimit = 3
	inputs[cvv].Width = 5
	inputs[cvv].Prompt = ""
	inputs[cvv].Validate = cvvValidator

	return PaymentCardModel{
		ctx:     ctx,
		client:  client,
		inputs:  inputs,
		focused: 0,
		err:     nil,
	}
}

func (m PaymentCardModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m PaymentCardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.inputs)-1 {
				return m, tea.Quit
			}
			m.nextInput()
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

	case errMsg:
		m.err = msg
		return m, nil
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m PaymentCardModel) View() string {
	return fmt.Sprintf(` %s
 %s

 %s  %s
 %s  %s

 %s
`,
		"Card Number",
		m.inputs[ccn].View(),
		"EXP",
		"CVV",
		m.inputs[exp].View(),
		m.inputs[cvv].View(),
		"Continue ->",
	) + "\n"
}

func (m *PaymentCardModel) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

func (m *PaymentCardModel) prevInput() {
	m.focused--
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}
