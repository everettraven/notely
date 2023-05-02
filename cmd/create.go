package cmd

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/everettraven/notely/pkg/noter"
	"github.com/everettraven/notely/pkg/texteditor"
	"github.com/spf13/cobra"
)

var createCmd = cobra.Command{
	Use:   "create name",
	Short: "create a new note",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		createNoter := noter.NewNoter()

		// fail early if a note with that name already exists
		// TODO(everettraven): add a prompt asking if the user would like to edit the note instead
		exists, err := createNoter.CheckNoteExists(args[0])
		if err != nil {
			log.Fatalf("error checking if note already exists: %s", err)
		} else if exists {
			log.Fatalf("note %q already exists", args[0])
		}

		p := tea.NewProgram(texteditor.NewTextEditor())

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
			err = createNoter.WriteNote(args[0], tedit.GetContents())
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}
