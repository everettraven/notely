package texteditor

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var _ tea.Model = &TextEditor{}

type TextEditor struct {
	textarea textarea.Model
	helper   help.Model
	err      error
	canceled bool
}

type TextEditorOption func(*TextEditor)

func WithContents(contents string) TextEditorOption {
	return func(t *TextEditor) {
		t.textarea.SetValue(contents)
	}
}

func NewTextEditor(opts ...TextEditorOption) *TextEditor {
	ta := textarea.New()
	ta.Focus()
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.CharLimit = 0
	ta.KeyMap.InsertNewline.SetEnabled(true)

	ta.SetHeight(10)
	ta.SetWidth(100)

	te := &TextEditor{
		textarea: ta,
		helper:   help.New(),
	}

	for _, opt := range opts {
		opt(te)
	}

	return te
}

func (t *TextEditor) GetContents() string {
	return t.textarea.Value()
}

func (t *TextEditor) Canceled() bool {
	return t.canceled
}

func (t *TextEditor) Init() tea.Cmd {
	return textarea.Blink
}

func (t *TextEditor) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var tCmd tea.Cmd
	t.textarea, tCmd = t.textarea.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return t, tea.Quit
		case tea.KeyCtrlQ:
			t.canceled = true
			return t, tea.Quit
		}

	// We handle errors just like any other message
	case error:
		t.err = msg
		return t, nil
	}

	return t, tCmd
}

func (t *TextEditor) View() string {
	sections := []string{}

	sections = append(sections, t.textarea.View())
	sections = append(sections, t.helper.View(t))

	return lipgloss.JoinVertical(0.2, sections...)
}

var saveExitBinding = key.NewBinding(
	key.WithKeys(tea.KeyEsc.String(), tea.KeyCtrlC.String()),
	key.WithHelp(tea.KeyCtrlC.String(), "Exit and save the note"),
	key.WithHelp(tea.KeyEsc.String(), "Exit and save the note"),
)
var cancelBinding = key.NewBinding(
	key.WithKeys(tea.KeyCtrlQ.String()),
	key.WithHelp(tea.KeyCtrlQ.String(), "Exit without saving"),
)

// ShortHelp returns bindings to show in the abbreviated help view. It's part
// of the help.KeyMap interface.
func (t *TextEditor) ShortHelp() []key.Binding {
	kb := []key.Binding{
		saveExitBinding,
		cancelBinding,
	}

	return kb
}

// FullHelp returns bindings to show the full help view. It's part of the
// help.KeyMap interface.
func (t *TextEditor) FullHelp() [][]key.Binding {
	kb := [][]key.Binding{{
		saveExitBinding,
		cancelBinding,
	}}

	return kb
}
