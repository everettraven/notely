package cmd

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/everettraven/notely/pkg/noter"
	"github.com/everettraven/notely/pkg/texteditor"
	"github.com/spf13/cobra"
)

var editCmd = cobra.Command{
	Use:   "edit name",
	Short: "edit an existing note",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		editNoter := noter.NewNoter()
		contents, err := editNoter.LoadNote(args[0])
		if err != nil {
			log.Fatal(err)
		}
		p := tea.NewProgram(texteditor.NewTextEditor(texteditor.WithContents(contents)))

		model, err := p.Run()
		if err != nil {
			log.Fatal(err)
		}

		tedit, ok := model.(*texteditor.TextEditor)
		if !ok {
			log.Fatal("model not expected type")
		}

		// Only write the note if a user has exited in a way that should save the note
		if !tedit.Canceled() {
			err = editNoter.WriteNote(args[0], tedit.GetContents())
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}
