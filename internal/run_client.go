package internal

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/SpaceSlow/gophkeeper/internal/application/models"
)

func RunClient() error {
	model := models.NewMainModel(context.Background(), "https://localhost/api/")
	_, err := tea.NewProgram(model, tea.WithAltScreen()).Run()
	return err
}
