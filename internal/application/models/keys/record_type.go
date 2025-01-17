package keys

import (
	"github.com/charmbracelet/bubbles/key"

	"github.com/SpaceSlow/gophkeeper/pkg/bubblekey"
)

type ChoiceFormKeyMap struct {
	PrevChoice key.Binding
	NextChoice key.Binding
	Select     key.Binding
	Back       key.Binding
	Quit       key.Binding
}

func (k ChoiceFormKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Back, k.Quit}
}

func (k ChoiceFormKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.PrevChoice, k.NextChoice, k.Back},
		{k.Select, bubblekey.Blank, k.Quit},
	}
}

var ChoiceFormKeys = ChoiceFormKeyMap{
	PrevChoice: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "up"),
	),
	NextChoice: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "down"),
	),
	Select: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
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
