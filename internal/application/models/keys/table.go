package keys

import "github.com/charmbracelet/bubbles/key"

type TableKeyMap struct {
	PrevLine key.Binding
	NextLine key.Binding
	Open     key.Binding
	Create   key.Binding
	Back     key.Binding
	Quit     key.Binding
}

func (k TableKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Back, k.Quit}
}

func (k TableKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.PrevLine, k.NextLine, k.Back},
		{k.Open, k.Create, k.Quit},
	}
}

var TableKeys = TableKeyMap{
	PrevLine: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "prev"),
	),
	NextLine: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "next"),
	),
	Open: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "open"),
	),
	Create: key.NewBinding(
		key.WithKeys("ctrl+n"),
		key.WithHelp("ctrl+n", "create"),
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
