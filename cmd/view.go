package cmd

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/everettraven/notely/pkg/noter"
	"github.com/everettraven/notely/pkg/viewer"
	"github.com/spf13/cobra"
)

var viewCmd = cobra.Command{
	Use:   "view name",
	Short: "view an existing note",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		viewNoter := noter.NewNoter()
		contents, err := viewNoter.LoadNote(args[0])
		if err != nil {
			log.Fatal(err)
		}

		view, err := viewer.NewViewer(contents)
		if err != nil {
			log.Fatal(err)
		}

		p := tea.NewProgram(view)

		_, err = p.Run()
		if err != nil {
			log.Fatal(err)
		}

	},
}
