package keys

import "github.com/charmbracelet/bubbles/key"

type BackQuitKeyMap struct {
	Back key.Binding
	Quit key.Binding
}

func (k BackQuitKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Back, k.Quit}
}

func (k BackQuitKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Back},
		{k.Quit},
	}
}

var BackQuitKeys = BackQuitKeyMap{
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
}
