package keys

import "github.com/charmbracelet/bubbles/key"

type RegisterKeyMap struct {
	PrevInput key.Binding
	NextInput key.Binding
	Enter     key.Binding
	Back      key.Binding
	Quit      key.Binding
}

func (k RegisterKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Back, k.Quit}
}

func (k RegisterKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.NextInput, k.PrevInput, k.Enter},
		{k.Back, k.Quit},
	}
}

var RegisterKeys = RegisterKeyMap{
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
		key.WithHelp("enter", "next/register"),
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
