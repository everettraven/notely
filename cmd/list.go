package cmd

import (
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/everettraven/notely/pkg/lister"
	"github.com/everettraven/notely/pkg/noter"
	"github.com/spf13/cobra"
)

var listCmd = cobra.Command{
	Use:   "list",
	Short: "list all notely notes",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		listNoter := noter.NewNoter()
		notes, err := listNoter.ListNotes()
		if err != nil {
			log.Fatalf("listing notes: %s", err)
		}

		items := []list.Item{}
		for _, note := range notes {
			items = append(items, lister.NewItem(note))
		}

		p := tea.NewProgram(lister.NewLister(items))

		_, err = p.Run()
		if err != nil {
			log.Fatal(err)
		}
	},
}
