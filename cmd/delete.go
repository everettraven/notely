package cmd

import (
	"github.com/everettraven/notely/pkg/noter"
	"github.com/spf13/cobra"
)

var deleteCmd = cobra.Command{
	Use:   "delete name",
	Short: "delete a note",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		deleteNoter := noter.NewNoter()
		deleteNoter.DeleteNote(args[0])
	},
}
