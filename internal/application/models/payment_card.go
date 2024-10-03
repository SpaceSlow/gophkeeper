package models

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"strconv"
	"strings"

	"github.com/SpaceSlow/gophkeeper/internal/domain/sensitive_records"
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
	cardholder
	metadata
)

type PaymentCardFormModel struct {
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

func NewPaymentCardFormModel(ctx context.Context, client *openapi.ClientWithResponses) tea.Model {
	var inputs = make([]textinput.Model, 5)
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

	inputs[cardholder] = textinput.New()
	inputs[cardholder].Placeholder = "Ivan Ivanov"
	inputs[cardholder].CharLimit = 20
	inputs[cardholder].Width = 30
	inputs[cardholder].Prompt = ""

	inputs[metadata] = textinput.New()
	inputs[metadata].Placeholder = "some metadata"
	inputs[metadata].CharLimit = 100
	inputs[metadata].Width = 50
	inputs[metadata].Prompt = ""

	return PaymentCardFormModel{
		ctx:     ctx,
		client:  client,
		inputs:  inputs,
		focused: 0,
		err:     nil,
	}
}

func (m PaymentCardFormModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m PaymentCardFormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				Metadata: m.inputs[metadata].Value(),
				Type:     openapi.PaymentCard,
			})
			var data bytes.Buffer
			enc := gob.NewEncoder(&data)
			exps := strings.Split(m.inputs[exp].Value(), "/")
			expMonth, _ := strconv.ParseUint(exps[0], 10, 64)
			expYear, _ := strconv.ParseUint(exps[1], 10, 64)
			code, _ := strconv.ParseInt(m.inputs[cvv].Value(), 10, 64)

			paymentCard := sensitive_records.PaymentCard{
				Number:     m.inputs[ccn].Value(),
				ExpMonth:   uint8(expMonth),
				ExpYear:    uint8(expYear),
				Cardholder: m.inputs[cardholder].Value(),
				Code:       int16(code),
			}
			enc.Encode(paymentCard)
			_, _ = m.client.PostSensitiveRecordDataWithBodyWithResponse(
				m.ctx,
				response.JSON201.Id,
				"application/octet-stream",
				&data,
			)
			return NewPaymentCardModel(m.ctx, m.client, &paymentCard, m.inputs[metadata].Value()), nil
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

func (m PaymentCardFormModel) View() string {
	return fmt.Sprintf(` %s
 %s

 %s	%s
 %s	%s

 %s
 %s

 %s: %s

 %s
`,
		"Card Number",
		m.inputs[ccn].View(),
		"EXP",
		"CVV",
		m.inputs[exp].View(),
		m.inputs[cvv].View(),
		"Cardholder",
		m.inputs[cardholder].View(),
		"Metadata",
		m.inputs[metadata].View(),
		"Continue ->",
	)
}

func (m *PaymentCardFormModel) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

func (m *PaymentCardFormModel) prevInput() {
	m.focused--
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}

type PaymentCardModel struct {
	ctx    context.Context
	client *openapi.ClientWithResponses

	paymentCard *sensitive_records.PaymentCard
	metadata    string
}

func NewPaymentCardModel(
	ctx context.Context,
	client *openapi.ClientWithResponses,
	paymentCard *sensitive_records.PaymentCard,
	metadata string,
) tea.Model {
	return PaymentCardModel{
		ctx:         ctx,
		client:      client,
		paymentCard: paymentCard,
		metadata:    metadata,
	}
}

func (m PaymentCardModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m PaymentCardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m PaymentCardModel) View() string {
	return fmt.Sprintf(` %s
 %s

 %s	%s
 %d/%d	%d

 %s
 %s

 %s: %s
`,
		"Card Number",
		m.paymentCard.Number,
		"EXP",
		"CVV",
		m.paymentCard.ExpMonth,
		m.paymentCard.ExpYear,
		m.paymentCard.Code,
		"Cardholder",
		m.paymentCard.Cardholder,
		"Metadata",
		m.metadata,
	)
}
