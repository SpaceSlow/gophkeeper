package keys

import (
	"github.com/charmbracelet/bubbles/key"
)

type BinaryKeyMap struct {
	Enter key.Binding
	Save  key.Binding
	Back  key.Binding
	Quit  key.Binding
}

func (k BinaryKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Back, k.Quit}
}

func (k BinaryKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Save, k.Back},
		{k.Enter, k.Quit},
	}
}

var BinaryKeys = BinaryKeyMap{
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "enter"),
	),
	Save: key.NewBinding(
		key.WithKeys("ctrl+s"),
		key.WithHelp("ctrl+s", "save file"),
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
