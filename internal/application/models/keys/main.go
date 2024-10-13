package keys

import "github.com/charmbracelet/bubbles/key"

type MainKeyMap struct {
	Register key.Binding
	Login    key.Binding
	Help     key.Binding
	Quit     key.Binding
}

func (k MainKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k MainKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Register, k.Login},
		{k.Help, k.Quit},
	}
}

var MainKeys = MainKeyMap{
	Register: key.NewBinding(
		key.WithKeys("1"),
		key.WithHelp("1", "register"),
	),
	Login: key.NewBinding(
		key.WithKeys("2"),
		key.WithHelp("2", "login"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("esc", "ctrl+c"),
		key.WithHelp("esc/ctrl+c", "quit"),
	),
}
