package lister

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

var _ tea.Model = &Lister{}

type Lister struct {
	list list.Model
	err  error
}

func NewLister(items []list.Item) *Lister {
	l := list.New(items, list.NewDefaultDelegate(), 15, 15)
	l.Title = "Notes"
	lister := &Lister{
		list: l,
	}

	return lister
}

func (l *Lister) Init() tea.Cmd {
	return nil
}

func (l *Lister) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var tCmd tea.Cmd
	l.list, tCmd = l.list.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return l, tea.Quit
		}

	// We handle errors just like any other message
	case error:
		l.err = msg
		return l, nil
	}

	return l, tCmd
}

func (l *Lister) View() string {
	return l.list.View()
}

type Item struct {
	title string
}

func NewItem(title string) Item {
	return Item{
		title: title,
	}
}

func (i Item) Title() string {
	return i.title
}

func (i Item) Description() string {
	return ""
}

func (i Item) FilterValue() string {
	return i.title
}
