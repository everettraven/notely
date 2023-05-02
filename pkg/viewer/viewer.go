package viewer

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
)

var _ tea.Model = &Viewer{}

type Viewer struct {
	viewport viewport.Model
	err      error
}

func NewViewer(contents string) (*Viewer, error) {
	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(100),
	)
	if err != nil {
		return nil, err
	}

	str, err := renderer.Render(contents)
	if err != nil {
		return nil, err
	}
	viewer := &Viewer{
		viewport: viewport.New(100, 20),
	}

	viewer.viewport.SetContent(str)

	return viewer, nil
}

func (v *Viewer) Init() tea.Cmd {
	return nil
}

func (v *Viewer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var tCmd tea.Cmd
	v.viewport, tCmd = v.viewport.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return v, tea.Quit
		}

	// We handle errors just like any other message
	case error:
		v.err = msg
		return v, nil
	}

	return v, tCmd
}

func (v *Viewer) View() string {
	return v.viewport.View()
}
