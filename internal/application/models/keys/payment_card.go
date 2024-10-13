package keys

import (
	"github.com/charmbracelet/bubbles/key"

	"github.com/SpaceSlow/gophkeeper/pkg/bubblekey"
)

type PaymentCardFormKeyMap struct {
	PrevInput key.Binding
	NextInput key.Binding
	Enter     key.Binding
	Back      key.Binding
	Quit      key.Binding
}

func (k PaymentCardFormKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Back, k.Quit}
}

func (k PaymentCardFormKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.PrevInput, k.NextInput, k.Back},
		{k.Enter, bubblekey.Blank, k.Quit},
	}
}

var PaymentCardFormKeys = PaymentCardFormKeyMap{
	PrevInput: key.NewBinding(
		key.WithKeys("up", "shift+tab"),
		key.WithHelp("↑/shift+tab", "prev"),
	),
	NextInput: key.NewBinding(
		key.WithKeys("down", "tab"),
		key.WithHelp("↓/tab", "next"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "next/send"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
}
